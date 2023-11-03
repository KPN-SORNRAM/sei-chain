package evmrpc_test

import (
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

func TestFilterNew(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name      string
		fromBlock string
		toBlock   string
		blockHash common.Hash
		addrs     []common.Address
		topics    [][]common.Hash
		wantErr   bool
	}{
		{
			name:      "happy path",
			fromBlock: "0x1",
			toBlock:   "0x2",
			addrs:     []common.Address{common.HexToAddress("0x123")},
			topics:    [][]common.Hash{{common.HexToHash("0x456")}},
			wantErr:   false,
		},
		{
			name:      "error: block hash and block range both given",
			fromBlock: "0x1",
			toBlock:   "0x2",
			blockHash: common.HexToHash("0xabc"),
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filterCriteria := map[string]interface{}{
				"fromBlock": tt.fromBlock,
				"toBlock":   tt.toBlock,
				"address":   tt.addrs,
				"topics":    tt.topics,
			}
			if tt.blockHash != (common.Hash{}) {
				filterCriteria["blockHash"] = tt.blockHash.Hex()
			}
			if len(tt.fromBlock) > 0 || len(tt.toBlock) > 0 {
				filterCriteria["fromBlock"] = tt.fromBlock
				filterCriteria["toBlock"] = tt.toBlock
			}
			resObj := sendRequestGood(t, "newFilter", filterCriteria)
			if tt.wantErr {
				_, ok := resObj["error"]
				require.True(t, ok)
			} else {
				got := resObj["result"].(float64)
				// make sure next filter id is not equal to this one
				resObj := sendRequestGood(t, "newFilter", filterCriteria)
				got2 := resObj["result"].(float64)
				require.NotEqual(t, got, got2)
			}
		})
	}
}

func TestFilterUninstall(t *testing.T) {
	t.Parallel()
	// uninstall existing filter
	filterCriteria := map[string]interface{}{
		"fromBlock": "0x1",
		"toBlock":   "0xa",
	}
	resObj := sendRequestGood(t, "newFilter", filterCriteria)
	filterId := int(resObj["result"].(float64))
	require.GreaterOrEqual(t, filterId, 1)

	resObj = sendRequest(t, TestPort, "uninstallFilter", filterId)
	uninstallSuccess := resObj["result"].(bool)
	require.True(t, uninstallSuccess)

	// uninstall non-existing filter
	nonExistingFilterId := 100
	resObj = sendRequest(t, TestPort, "uninstallFilter", nonExistingFilterId)
	uninstallSuccess = resObj["result"].(bool)
	require.False(t, uninstallSuccess)
}

func TestFilterGetLogs(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name      string
		blockHash common.Hash
		fromBlock string
		toBlock   string
		addrs     []common.Address
		topics    [][]common.Hash
		wantErr   bool
		wantLen   int
		check     func(t *testing.T, log map[string]interface{})
	}{
		{
			name:      "filter by single address",
			fromBlock: "0x2",
			toBlock:   "0x2",
			addrs:     []common.Address{common.HexToAddress("0x1111111111111111111111111111111111111112")},
			wantErr:   false,
			check: func(t *testing.T, log map[string]interface{}) {
				require.Equal(t, "0x1111111111111111111111111111111111111112", log["address"].(string))
			},
			wantLen: 1,
		},
		{
			name:      "filter by single topic",
			blockHash: common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000"),
			fromBlock: "0x3",
			toBlock:   "0x3",
			topics:    [][]common.Hash{{common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000123")}},
			wantErr:   false,
			check: func(t *testing.T, log map[string]interface{}) {
				require.Equal(t, "0x0000000000000000000000000000000000000000000000000000000000000123", log["topics"].([]interface{})[0].(string))
			},
			wantLen: 1,
		},
		{
			name:      "multiple addresses, multiple topics",
			blockHash: common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000"),
			fromBlock: "0x2",
			toBlock:   "0x2",
			addrs: []common.Address{
				common.HexToAddress("0x1111111111111111111111111111111111111112"),
				common.HexToAddress("0x1111111111111111111111111111111111111113"),
			},
			topics: [][]common.Hash{
				{common.Hash(common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000123"))},
				{common.Hash(common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000456"))},
			},
			wantErr: false,
			check: func(t *testing.T, log map[string]interface{}) {
				if log["address"].(string) != "0x1111111111111111111111111111111111111112" && log["address"].(string) != "0x1111111111111111111111111111111111111113" {
					t.Fatalf("address %s not in expected list", log["address"].(string))
				}
				firstTopic := log["topics"].([]interface{})[0].(string)
				secondTopic := log["topics"].([]interface{})[1].(string)
				require.Equal(t, "0x0000000000000000000000000000000000000000000000000000000000000123", firstTopic)
				require.Equal(t, "0x0000000000000000000000000000000000000000000000000000000000000456", secondTopic)
			},
			wantLen: 2,
		},
		{
			name:      "wildcard first topic",
			blockHash: common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000"),
			fromBlock: "0x2",
			toBlock:   "0x2",
			topics: [][]common.Hash{
				{},
				{common.Hash(common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000456"))},
			},
			wantErr: false,
			check: func(t *testing.T, log map[string]interface{}) {
				secondTopic := log["topics"].([]interface{})[1].(string)
				require.Equal(t, "0x0000000000000000000000000000000000000000000000000000000000000456", secondTopic)
			},
			wantLen: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filterCriteria := map[string]interface{}{
				"fromBlock": tt.fromBlock,
				"toBlock":   tt.toBlock,
				"address":   tt.addrs,
				"topics":    tt.topics,
			}
			if tt.blockHash != (common.Hash{}) {
				filterCriteria["blockHash"] = tt.blockHash.Hex()
			}
			if len(tt.fromBlock) > 0 || len(tt.toBlock) > 0 {
				filterCriteria["fromBlock"] = tt.fromBlock
				filterCriteria["toBlock"] = tt.toBlock
			}
			resObj := sendRequestGood(t, "getLogs", filterCriteria)
			if tt.wantErr {
				_, ok := resObj["error"]
				require.True(t, ok)
			} else {
				got := resObj["result"].([]interface{})
				for _, log := range got {
					logObj := log.(map[string]interface{})
					tt.check(t, logObj)
				}
				require.Equal(t, len(got), tt.wantLen)
			}
		})
	}
}

func TestFilterGetFilterLogs(t *testing.T) {
	t.Parallel()
	filterCriteria := map[string]interface{}{
		"fromBlock": "0x4",
		"toBlock":   "0x4",
	}
	resObj := sendRequestGood(t, "newFilter", filterCriteria)
	filterId := int(resObj["result"].(float64))

	resObj = sendRequest(t, TestPort, "getFilterLogs", filterId)
	logs := resObj["result"].([]interface{})
	require.Equal(t, 1, len(logs))
	for _, log := range logs {
		logObj := log.(map[string]interface{})
		require.Equal(t, "0x4", logObj["blockNumber"].(string))
	}

	// error: filter id does not exist
	nonexistentFilterId := 1000
	resObj = sendRequest(t, TestPort, "getFilterLogs", nonexistentFilterId)
	_, ok := resObj["error"]
	require.True(t, ok)
}

func TestFilterGetFilterChanges(t *testing.T) {
	t.Parallel()
	filterCriteria := map[string]interface{}{
		"fromBlock": "0x5",
	}
	resObj := sendRequest(t, TestPort, "newFilter", filterCriteria)
	filterId := int(resObj["result"].(float64))

	resObj = sendRequest(t, TestPort, "getFilterChanges", filterId)
	logs := resObj["result"].([]interface{})
	require.Equal(t, 1, len(logs))
	logObj := logs[0].(map[string]interface{})
	require.Equal(t, "0x5", logObj["blockNumber"].(string))

	// another query
	resObj = sendRequest(t, TestPort, "getFilterChanges", filterId)
	logs = resObj["result"].([]interface{})
	require.Equal(t, 1, len(logs))
	logObj = logs[0].(map[string]interface{})
	require.Equal(t, "0x6", logObj["blockNumber"].(string))

	// error: filter id does not exist
	nonExistingFilterId := 1000
	resObj = sendRequest(t, TestPort, "getFilterChanges", nonExistingFilterId)
	_, ok := resObj["error"]
	require.True(t, ok)
}

func TestFilterExpiration(t *testing.T) {
	t.Parallel()
	filterCriteria := map[string]interface{}{
		"fromBlock": "0x1",
		"toBlock":   "0xa",
	}
	resObj := sendRequestGood(t, "newFilter", filterCriteria)
	filterId := int(resObj["result"].(float64))

	// wait for filter to expire
	time.Sleep(2 * filterTimeoutDuration)

	resObj = sendRequest(t, TestPort, "getFilterLogs", filterId)
	_, ok := resObj["error"]
	require.True(t, ok)
}

func TestFilterGetFilterLogsKeepsFilterAlive(t *testing.T) {
	t.Parallel()
	filterCriteria := map[string]interface{}{
		"fromBlock": "0x1",
		"toBlock":   "0xa",
	}
	resObj := sendRequestGood(t, "newFilter", filterCriteria)
	filterId := int(resObj["result"].(float64))

	for i := 0; i < 5; i++ {
		// should keep filter alive
		resObj = sendRequestGood(t, "getFilterLogs", filterId)
		_, ok := resObj["error"]
		require.False(t, ok)
		time.Sleep(filterTimeoutDuration / 2)
	}
}

func TestFilterGetFilterChangesKeepsFilterAlive(t *testing.T) {
	t.Parallel()
	filterCriteria := map[string]interface{}{
		"fromBlock": "0x1",
		"toBlock":   "0xa",
	}
	resObj := sendRequestGood(t, "newFilter", filterCriteria)
	filterId := int(resObj["result"].(float64))

	for i := 0; i < 5; i++ {
		// should keep filter alive
		resObj = sendRequestGood(t, "getFilterChanges", filterId)
		_, ok := resObj["error"]
		require.False(t, ok)
		time.Sleep(filterTimeoutDuration / 2)
	}
}
