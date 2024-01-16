package evmrpc

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"sync"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/eth/filters"
	"github.com/sei-protocol/sei-chain/utils"
	"github.com/sei-protocol/sei-chain/x/evm/keeper"
	rpcclient "github.com/tendermint/tendermint/rpc/client"
	"github.com/tendermint/tendermint/rpc/coretypes"
	tmtypes "github.com/tendermint/tendermint/types"
)

const TxSearchPerPage = 10

type FilterType byte

const (
	UnknownSubscription FilterType = iota
	LogsSubscription
	BlocksSubscription
)

type filter struct {
	typ      FilterType
	fc       filters.FilterCriteria
	deadline *time.Timer

	// BlocksSubscription
	blockCursor string

	// LogsSubscription
	lastToHeight int64
}

type FilterAPI struct {
	tmClient     rpcclient.Client
	nextFilterID uint64
	filtersMu    sync.Mutex
	filters      map[uint64]filter
	filterConfig *FilterConfig
	logFetcher   *LogFetcher
}

type FilterConfig struct {
	timeout time.Duration
}

type EventItemDataWrapper struct {
	Type  string          `json:"type"`
	Value json.RawMessage `json:"value"`
}

func NewFilterAPI(tmClient rpcclient.Client, logFetcher *LogFetcher, filterConfig *FilterConfig) *FilterAPI {
	filters := make(map[uint64]filter)
	api := &FilterAPI{
		tmClient:     tmClient,
		nextFilterID: 1,
		filtersMu:    sync.Mutex{},
		filters:      filters,
		filterConfig: filterConfig,
		logFetcher:   logFetcher,
	}

	go api.timeoutLoop(filterConfig.timeout)

	return api
}

func (a *FilterAPI) timeoutLoop(timeout time.Duration) {
	ticker := time.NewTicker(timeout)
	defer ticker.Stop()
	for {
		<-ticker.C
		a.filtersMu.Lock()
		for id, filter := range a.filters {
			select {
			case <-filter.deadline.C:
				delete(a.filters, id)
			default:
				continue
			}
		}
		a.filtersMu.Unlock()
	}
}

func (a *FilterAPI) NewFilter(
	_ context.Context,
	crit filters.FilterCriteria,
) (*uint64, error) {
	a.filtersMu.Lock()
	defer a.filtersMu.Unlock()
	curFilterID := a.nextFilterID
	a.nextFilterID++
	a.filters[curFilterID] = filter{
		typ:          LogsSubscription,
		fc:           crit,
		deadline:     time.NewTimer(a.filterConfig.timeout),
		lastToHeight: 0,
	}
	return &curFilterID, nil
}

func (a *FilterAPI) NewBlockFilter(
	_ context.Context,
) (*uint64, error) {
	a.filtersMu.Lock()
	defer a.filtersMu.Unlock()
	curFilterID := a.nextFilterID
	a.nextFilterID++
	a.filters[curFilterID] = filter{
		typ:         BlocksSubscription,
		deadline:    time.NewTimer(a.filterConfig.timeout),
		blockCursor: "",
	}
	return &curFilterID, nil
}

func (a *FilterAPI) GetFilterChanges(
	ctx context.Context,
	filterID uint64,
) (interface{}, error) {
	a.filtersMu.Lock()
	defer a.filtersMu.Unlock()
	filter, ok := a.filters[filterID]
	if !ok {
		return nil, errors.New("filter does not exist")
	}

	if !filter.deadline.Stop() {
		// timer expired but filter is not yet removed in timeout loop
		// receive timer value and reset timer
		<-filter.deadline.C
	}
	filter.deadline.Reset(a.filterConfig.timeout)

	switch filter.typ {
	case BlocksSubscription:
		hashes, cursor, err := a.getBlockHeadersAfter(ctx, filter.blockCursor)
		if err != nil {
			return nil, err
		}
		updatedFilter := a.filters[filterID]
		updatedFilter.blockCursor = cursor
		a.filters[filterID] = updatedFilter
		return hashes, nil
	case LogsSubscription:
		// filter by hash would have no updates if it has previously queried for this crit
		if filter.fc.BlockHash != nil && filter.lastToHeight > 0 {
			return nil, nil
		}
		// filter with a ToBlock would have no updates if it has previously queried for this crit
		if filter.fc.ToBlock != nil && filter.lastToHeight >= filter.fc.ToBlock.Int64() {
			return nil, nil
		}
		logs, lastToHeight, err := a.logFetcher.GetLogsByFilters(ctx, filter.fc, filter.lastToHeight)
		if err != nil {
			return nil, err
		}
		updatedFilter := a.filters[filterID]
		updatedFilter.lastToHeight = lastToHeight + 1
		a.filters[filterID] = updatedFilter
		return logs, nil
	default:
		return nil, errors.New("unknown filter type")
	}
}

func (a *FilterAPI) GetFilterLogs(
	ctx context.Context,
	filterID uint64,
) ([]*ethtypes.Log, error) {
	a.filtersMu.Lock()
	defer a.filtersMu.Unlock()
	filter, ok := a.filters[filterID]
	if !ok {
		return nil, errors.New("filter does not exist")
	}

	if !filter.deadline.Stop() {
		// timer expired but filter is not yet removed in timeout loop
		// receive timer value and reset timer
		<-filter.deadline.C
	}
	filter.deadline.Reset(a.filterConfig.timeout)

	logs, lastToHeight, err := a.logFetcher.GetLogsByFilters(ctx, filter.fc, 0)
	if err != nil {
		return nil, err
	}
	updatedFilter := a.filters[filterID]
	updatedFilter.lastToHeight = lastToHeight
	a.filters[filterID] = updatedFilter
	return logs, nil
}

func (a *FilterAPI) GetLogs(
	ctx context.Context,
	crit filters.FilterCriteria,
) ([]*ethtypes.Log, error) {
	logs, _, err := a.logFetcher.GetLogsByFilters(ctx, crit, 0)
	return logs, err
}

// get block headers after a certain cursor. Can use an empty string cursor
// to get the latest block header.
func (a *FilterAPI) getBlockHeadersAfter(
	ctx context.Context,
	cursor string,
) ([]common.Hash, string, error) {
	q := NewBlockQueryBuilder()
	builtQuery := q.Build()
	hasMore := true
	headers := []common.Hash{}
	for hasMore {
		res, err := a.tmClient.Events(ctx, &coretypes.RequestEvents{
			Filter: &coretypes.EventFilter{Query: builtQuery},
			After:  cursor,
		})
		if err != nil {
			return nil, "", err
		}
		hasMore = res.More
		cursor = res.Newest

		for _, item := range res.Items {
			wrapper := EventItemDataWrapper{}
			err := json.Unmarshal(item.Data, &wrapper)
			if err != nil {
				return nil, "", err
			}
			block := tmtypes.EventDataNewBlock{}
			err = json.Unmarshal(wrapper.Value, &block)
			if err != nil {
				return nil, "", err
			}
			headers = append(headers, common.BytesToHash(block.Block.Hash()))
		}
	}
	return headers, cursor, nil
}

func (a *FilterAPI) UninstallFilter(
	_ context.Context,
	filterID uint64,
) bool {
	a.filtersMu.Lock()
	defer a.filtersMu.Unlock()
	_, found := a.filters[filterID]
	if !found {
		return false
	}
	delete(a.filters, filterID)
	return true
}

type LogFetcher struct {
	tmClient    rpcclient.Client
	k           *keeper.Keeper
	ctxProvider func(int64) sdk.Context
}

func (f *LogFetcher) GetLogsByFilters(ctx context.Context, crit filters.FilterCriteria, lastToHeight int64) ([]*ethtypes.Log, int64, error) {
	bloomIndexes := EncodeFilters(crit.Addresses, crit.Topics)
	if crit.BlockHash != nil {
		block, err := blockByHashWithRetry(ctx, f.tmClient, crit.BlockHash[:])
		if err != nil {
			return nil, 0, err
		}
		return f.GetLogsForBlock(ctx, block, crit, bloomIndexes), block.Block.Height, nil
	}
	latest := f.ctxProvider(LatestCtxHeight).BlockHeight()
	begin, end := int64(0), latest
	if crit.FromBlock != nil {
		begin = getHeightFromBigIntBlockNumber(latest, crit.FromBlock)
	}
	if crit.ToBlock != nil {
		end = getHeightFromBigIntBlockNumber(latest, crit.ToBlock)
	}
	if lastToHeight > begin {
		begin = lastToHeight
	}
	blockHeights := f.FindBlockesByBloom(begin, end, bloomIndexes)
	res := []*ethtypes.Log{}
	for _, height := range blockHeights {
		h := height
		block, err := blockWithRetry(ctx, f.tmClient, &h)
		if err != nil {
			return nil, 0, err
		}
		res = append(res, f.GetLogsForBlock(ctx, block, crit, bloomIndexes)...)
	}

	return res, end, nil
}

func (f *LogFetcher) GetLogsForBlock(ctx context.Context, block *coretypes.ResultBlock, crit filters.FilterCriteria, filters [][]bloomIndexes) []*ethtypes.Log {
	possibleLogs := f.FindLogsByBloom(block.Block.Height, filters)
	matchedLogs := utils.Filter(possibleLogs, func(l *ethtypes.Log) bool { return f.IsLogExactMatch(l, crit) })
	for i, l := range matchedLogs {
		l.BlockHash = common.Hash(block.BlockID.Hash)
		l.Index = uint(i)
	}
	return matchedLogs
}

func (f *LogFetcher) FindBlockesByBloom(begin, end int64, filters [][]bloomIndexes) (res []int64) {
	//TODO: parallelize
	ctx := f.ctxProvider(LatestCtxHeight)
	for height := begin; height <= end; height++ {
		blockBloom := f.k.GetBlockBloom(ctx, height)
		if MatchFilters(blockBloom, filters) {
			res = append(res, height)
		}
	}
	return
}

func (f *LogFetcher) FindLogsByBloom(height int64, filters [][]bloomIndexes) (res []*ethtypes.Log) {
	ctx := f.ctxProvider(LatestCtxHeight)
	txHashes := f.k.GetTxHashesOnHeight(ctx, height)
	for _, hash := range txHashes {
		receipt, err := f.k.GetReceipt(ctx, hash)
		if err != nil {
			ctx.Logger().Error(fmt.Sprintf("FindLogsByBloom: unable to find receipt for hash %s", hash.Hex()))
			continue
		}
		if MatchFilters(ethtypes.Bloom(receipt.LogsBloom), filters) {
			res = append(res, keeper.GetLogsForTx(receipt)...)
		}
	}
	return
}

func (f *LogFetcher) IsLogExactMatch(log *ethtypes.Log, crit filters.FilterCriteria) bool {
	addrMatch := len(crit.Addresses) == 0
	for _, addrFilter := range crit.Addresses {
		if log.Address == addrFilter {
			addrMatch = true
			break
		}
	}
	return addrMatch && matchTopics(crit.Topics, log.Topics)
}

func matchTopics(topics [][]common.Hash, eventTopics []common.Hash) bool {
	for i, topicList := range topics {
		if len(topicList) == 0 {
			// anything matches for this position
			continue
		}
		if i >= len(eventTopics) {
			return false
		}
		matched := false
		for _, topic := range topicList {
			if topic == eventTopics[i] {
				matched = true
				break
			}
		}
		if !matched {
			return false
		}
	}
	return true
}
