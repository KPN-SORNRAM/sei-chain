package state

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/sei-protocol/sei-chain/utils"
	"github.com/sei-protocol/sei-chain/x/evm/types"
)

// Initialized for each transaction individually
type DBImpl struct {
	ctx             sdk.Context
	snapshottedCtxs []sdk.Context

	logs            []*ethtypes.Log
	snapshottedLogs [][]*ethtypes.Log
	// If err is not nil at the end of the execution, the transaction will be rolled
	// back.
	err error

	// a temporary address that facilitates transfer during the processing of this particular
	// transaction. Its account state and balance will be deleted before the block commits
	middleManAddress sdk.AccAddress
	// a temporary address that collects fees for this particular transaction so that there is
	// no single bottleneck for fee collection. Its account state and balance will be deleted
	// before the block commits
	coinbaseAddress sdk.AccAddress

	k          EVMKeeper
	simulation bool

	transientStates       map[string]map[string]common.Hash
	transientAccounts     map[string][]byte
	transientModuleStates map[string][]byte
}

func NewDBImpl(ctx sdk.Context, k EVMKeeper, simulation bool) *DBImpl {
	s := &DBImpl{
		ctx:                   ctx,
		k:                     k,
		snapshottedCtxs:       []sdk.Context{},
		middleManAddress:      GetMiddleManAddress(ctx),
		coinbaseAddress:       GetCoinbaseAddress(ctx),
		simulation:            simulation,
		transientStates:       make(map[string]map[string]common.Hash),
		transientAccounts:     make(map[string][]byte),
		transientModuleStates: make(map[string][]byte),
	}
	s.Snapshot() // take an initial snapshot for GetCommitted
	return s
}

func (s *DBImpl) SetEVM(evm *vm.EVM) {
	s.ctx = types.SetCtxEVM(s.ctx, evm)
	s.snapshottedCtxs = utils.Map(s.snapshottedCtxs, func(ctx sdk.Context) sdk.Context { return types.SetCtxEVM(ctx, evm) })
}

// AddPreimage records a SHA3 preimage seen by the VM.
// AddPreimage performs a no-op since the EnablePreimageRecording flag is disabled
// on the vm.Config during state transitions. No store trie preimages are written
// to the database.
func (s *DBImpl) AddPreimage(_ common.Hash, _ []byte) {}

func (s *DBImpl) Finalize() error {
	if s.simulation {
		panic("should never call finalize on a simulation DB")
	}
	if s.err != nil {
		return s.err
	}

	// remove transient states
	// write cache to underlying
	s.flushCtx(s.ctx)
	// write all snapshotted caches in reverse order, except the very first one (base) which will be written by baseapp::runTx
	for i := len(s.snapshottedCtxs) - 1; i > 0; i-- {
		s.flushCtx(s.snapshottedCtxs[i])
	}
	return nil
}

func (s *DBImpl) flushCtx(ctx sdk.Context) {
	ctx.MultiStore().(sdk.CacheMultiStore).Write()
}

// Backward-compatibility functions
func (s *DBImpl) Error() error {
	return s.Err()
}

func (s *DBImpl) GetStorageRoot(common.Address) common.Hash {
	panic("GetStorageRoot is not implemented and called unexpectedly")
}

func (s *DBImpl) Copy() vm.StateDB {
	newCtx := s.ctx.WithMultiStore(s.ctx.MultiStore().CacheMultiStore())
	return &DBImpl{
		ctx:                   newCtx,
		snapshottedCtxs:       append(s.snapshottedCtxs, s.ctx),
		logs:                  []*ethtypes.Log{},
		snapshottedLogs:       append(s.snapshottedLogs, s.logs),
		k:                     s.k,
		middleManAddress:      s.middleManAddress,
		coinbaseAddress:       s.coinbaseAddress,
		simulation:            s.simulation,
		err:                   s.err,
		transientStates:       s.transientStates,
		transientAccounts:     s.transientAccounts,
		transientModuleStates: s.transientModuleStates,
	}
}

func (s *DBImpl) Finalise(bool) {
	s.ctx.Logger().Info("Finalise should only be called during simulation and will no-op")
}

func (s *DBImpl) Commit(uint64, bool) (common.Hash, error) {
	panic("Commit is not implemented and called unexpectedly")
}

func (s *DBImpl) SetTxContext(common.Hash, int) {
	//noop
}

func (s *DBImpl) IntermediateRoot(bool) common.Hash {
	panic("IntermediateRoot is not implemented and called unexpectedly")
}

func (s *DBImpl) TxIndex() int {
	return s.ctx.TxIndex()
}

func (s *DBImpl) Preimages() map[common.Hash][]byte {
	return map[common.Hash][]byte{}
}

// ** TEST ONLY FUNCTIONS **//
func (s *DBImpl) Err() error {
	return s.err
}

func (s *DBImpl) WithErr(err error) {
	s.err = err
}

func (s *DBImpl) Ctx() sdk.Context {
	return s.ctx
}

func (s *DBImpl) WithCtx(ctx sdk.Context) {
	s.ctx = ctx
}
