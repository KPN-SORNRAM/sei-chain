package ibc_test

import (
	"errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	clienttypes "github.com/cosmos/ibc-go/v3/modules/core/02-client/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
	pcommon "github.com/sei-protocol/sei-chain/precompiles/common"
	"github.com/sei-protocol/sei-chain/precompiles/ibc"
	testkeeper "github.com/sei-protocol/sei-chain/testutil/keeper"
	"github.com/sei-protocol/sei-chain/x/evm/state"
	"github.com/stretchr/testify/require"
	tmtypes "github.com/tendermint/tendermint/proto/tendermint/types"
	"math/big"
	"reflect"
	"testing"
)

type MockTransferKeeper struct{}

func (tk *MockTransferKeeper) SendTransfer(ctx sdk.Context, sourcePort, sourceChannel string, token sdk.Coin,
	sender sdk.AccAddress, receiver string, timeoutHeight clienttypes.Height, timeoutTimestamp uint64) error {
	return nil
}

type MockFailedTransferTransferKeeper struct{}

func (tk *MockFailedTransferTransferKeeper) SendTransfer(ctx sdk.Context, sourcePort, sourceChannel string, token sdk.Coin,
	sender sdk.AccAddress, receiver string, timeoutHeight clienttypes.Height, timeoutTimestamp uint64) error {
	return errors.New("failed to send transfer")
}

func TestPrecompile_Run(t *testing.T) {
	senderSeiAddress, senderEvmAddress := testkeeper.MockAddressPair()
	_, receiverEvmAddress := testkeeper.MockAddressPair()

	pre, _ := ibc.NewPrecompile(nil, nil)
	testTransfer, _ := pre.ABI.MethodById(pre.TransferID)
	packedTrue, _ := testTransfer.Outputs.Pack(true)

	type fields struct {
		transferKeeper pcommon.TransferKeeper
	}

	type input struct {
		receiverEvmAddr  common.Address
		sourcePort       string
		sourceChannel    string
		denom            string
		amount           *big.Int
		revisionNumber   uint64
		revisionHeight   uint64
		timeoutTimestamp uint64
	}
	type args struct {
		caller      common.Address
		input       *input
		suppliedGas uint64
		value       *big.Int
	}

	commonArgs := args{
		caller: senderEvmAddress,
		input: &input{
			receiverEvmAddr:  receiverEvmAddress,
			sourcePort:       "sourcePort",
			sourceChannel:    "sourceChannel",
			denom:            "denom",
			amount:           big.NewInt(100),
			revisionNumber:   1,
			revisionHeight:   1,
			timeoutTimestamp: 1,
		},
		suppliedGas: uint64(1000000),
		value:       nil,
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantBz  []byte
		wantErr bool
	}{
		{
			name:    "successful transfer: with amount > 0 between EVM addresses",
			fields:  fields{transferKeeper: &MockTransferKeeper{}},
			args:    commonArgs,
			wantBz:  packedTrue,
			wantErr: false,
		},
		{
			name:    "failed transfer: internal error",
			fields:  fields{transferKeeper: &MockFailedTransferTransferKeeper{}},
			args:    commonArgs,
			wantBz:  nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testApp := testkeeper.EVMTestApp
			ctx := testApp.NewContext(false, tmtypes.Header{}).WithBlockHeight(2)
			k := &testApp.EvmKeeper
			k.SetAddressMapping(ctx, senderSeiAddress, senderEvmAddress)
			stateDb := state.NewDBImpl(ctx, k, true)
			evm := vm.EVM{
				StateDB:   stateDb,
				TxContext: vm.TxContext{Origin: senderEvmAddress},
			}
			p, _ := ibc.NewPrecompile(tt.fields.transferKeeper, k)
			transfer, err := p.ABI.MethodById(p.TransferID)
			require.Nil(t, err)
			inputs, err := transfer.Inputs.Pack(tt.args.input.receiverEvmAddr,
				tt.args.input.sourcePort, tt.args.input.sourceChannel, tt.args.input.denom, tt.args.input.amount,
				tt.args.input.revisionNumber, tt.args.input.revisionHeight, tt.args.input.timeoutTimestamp)
			require.Nil(t, err)
			gotBz, _, err := p.RunAndCalculateGas(&evm, tt.args.caller, common.Address{}, append(p.TransferID, inputs...), tt.args.suppliedGas, tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("Run() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotBz, tt.wantBz) {
				t.Errorf("Run() gotBz = %v, want %v", gotBz, tt.wantBz)
			}
		})
	}
}
