package wasmd_test

import (
	"fmt"
	"os"
	"testing"
	"time"

	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/sei-protocol/sei-chain/app"
	"github.com/sei-protocol/sei-chain/precompiles/wasmd"
	testkeeper "github.com/sei-protocol/sei-chain/testutil/keeper"
	"github.com/sei-protocol/sei-chain/x/evm/state"
	"github.com/stretchr/testify/require"
)

func TestRequiredGas(t *testing.T) {
	testApp := app.Setup(false, false)
	p, err := wasmd.NewPrecompile(&testApp.EvmKeeper, wasmkeeper.NewDefaultPermissionKeeper(testApp.WasmKeeper), testApp.WasmKeeper)
	require.Nil(t, err)
	require.Equal(t, uint64(2000), p.RequiredGas(p.ExecuteID))
	require.Equal(t, uint64(2000), p.RequiredGas(p.InstantiateID))
	require.Equal(t, uint64(1000), p.RequiredGas(p.QueryID))
	require.Equal(t, uint64(0), p.RequiredGas([]byte{15, 15, 15, 15})) // invalid method
}

func TestAddress(t *testing.T) {
	testApp := app.Setup(false, false)
	p, err := wasmd.NewPrecompile(&testApp.EvmKeeper, wasmkeeper.NewDefaultPermissionKeeper(testApp.WasmKeeper), testApp.WasmKeeper)
	require.Nil(t, err)
	require.Equal(t, "0x0000000000000000000000000000000000001002", p.Address().Hex())
}

func TestInstantiate(t *testing.T) {
	testApp := app.Setup(false, false)
	mockAddr, mockEVMAddr := testkeeper.MockAddressPair()
	ctx := testApp.GetContextForDeliverTx([]byte{}).WithBlockTime(time.Now())
	wasmKeeper := wasmkeeper.NewDefaultPermissionKeeper(testApp.WasmKeeper)
	p, err := wasmd.NewPrecompile(&testApp.EvmKeeper, wasmKeeper, testApp.WasmKeeper)
	require.Nil(t, err)
	code, err := os.ReadFile("../../example/cosmwasm/echo/artifacts/echo.wasm")
	require.Nil(t, err)
	codeID, err := wasmKeeper.Create(ctx, mockAddr, code, nil)
	require.Nil(t, err)
	instantiateMethod, err := p.ABI.MethodById(p.InstantiateID)
	require.Nil(t, err)
	amtsbz, err := sdk.NewCoins().MarshalJSON()
	require.Nil(t, err)
	args, err := instantiateMethod.Inputs.Pack(
		codeID,
		mockAddr.String(),
		[]byte("{}"),
		"test",
		amtsbz,
	)
	statedb := state.NewDBImpl(ctx, &testApp.EvmKeeper, true)
	evm := vm.EVM{
		StateDB: statedb,
	}
	res, err := p.Run(&evm, mockEVMAddr, append(p.InstantiateID, args...))
	require.Nil(t, err)
	outputs, err := instantiateMethod.Outputs.Unpack(res)
	require.Nil(t, err)
	require.Equal(t, 2, len(outputs))
	require.Equal(t, "sei14hj2tavq8fpesdwxxcu44rty3hh90vhujrvcmstl4zr3txmfvw9sh9m79m", outputs[0].(string))
	require.Empty(t, outputs[1].([]byte))

	// non-existent code ID
	args, _ = instantiateMethod.Inputs.Pack(
		codeID+1,
		mockAddr.String(),
		[]byte("{}"),
		"test",
		amtsbz,
	)
	res, err = p.Run(&evm, mockEVMAddr, append(p.InstantiateID, args...))
	require.NotNil(t, err)

	// bad inputs
	badArgs, _ := instantiateMethod.Inputs.Pack(codeID, "not bech32", []byte("{}"), "test", amtsbz)
	_, err = p.Run(&evm, mockEVMAddr, append(p.InstantiateID, badArgs...))
	require.NotNil(t, err)
	badArgs, _ = instantiateMethod.Inputs.Pack(codeID, mockAddr.String(), []byte("{}"), "test", []byte("bad coins"))
	_, err = p.Run(&evm, mockEVMAddr, append(p.InstantiateID, badArgs...))
	require.NotNil(t, err)
}

func TestExecute(t *testing.T) {
	testApp := app.Setup(false, false)
	mockAddr, mockEVMAddr := testkeeper.MockAddressPair()
	ctx := testApp.GetContextForDeliverTx([]byte{}).WithBlockTime(time.Now())
	testApp.EvmKeeper.SetAddressMapping(ctx, mockAddr, mockEVMAddr)
	wasmKeeper := wasmkeeper.NewDefaultPermissionKeeper(testApp.WasmKeeper)
	p, err := wasmd.NewPrecompile(&testApp.EvmKeeper, wasmKeeper, testApp.WasmKeeper)
	require.Nil(t, err)
	code, err := os.ReadFile("../../example/cosmwasm/echo/artifacts/echo.wasm")
	require.Nil(t, err)
	codeID, err := wasmKeeper.Create(ctx, mockAddr, code, nil)
	require.Nil(t, err)
	contractAddr, _, err := wasmKeeper.Instantiate(ctx, codeID, mockAddr, mockAddr, []byte("{}"), "test", sdk.NewCoins())
	require.Nil(t, err)

	amts := sdk.NewCoins(sdk.NewCoin("usei", sdk.NewInt(1000)))
	testApp.BankKeeper.MintCoins(ctx, "evm", amts)
	testApp.BankKeeper.SendCoinsFromModuleToAccount(ctx, "evm", mockAddr, amts)
	amtsbz, err := amts.MarshalJSON()
	require.Nil(t, err)
	executeMethod, err := p.ABI.MethodById(p.ExecuteID)
	require.Nil(t, err)
	args, err := executeMethod.Inputs.Pack(contractAddr.String(), []byte("{\"echo\":{\"message\":\"test msg\"}}"), amtsbz)
	require.Nil(t, err)
	statedb := state.NewDBImpl(ctx, &testApp.EvmKeeper, true)
	evm := vm.EVM{
		StateDB: statedb,
	}
	res, err := p.Run(&evm, mockEVMAddr, append(p.ExecuteID, args...))
	require.Nil(t, err)
	outputs, err := executeMethod.Outputs.Unpack(res)
	require.Nil(t, err)
	require.Equal(t, 1, len(outputs))
	require.Equal(t, fmt.Sprintf("received test msg from %s with 1000usei", mockAddr.String()), string(outputs[0].([]byte)))

	// bad contract address
	args, _ = executeMethod.Inputs.Pack(mockAddr.String(), []byte("{\"echo\":{\"message\":\"test msg\"}}"), amtsbz)
	res, err = p.Run(&evm, mockEVMAddr, append(p.ExecuteID, args...))
	require.NotNil(t, err)

	// bad inputs
	args, _ = executeMethod.Inputs.Pack("not bech32", []byte("{\"echo\":{\"message\":\"test msg\"}}"), amtsbz)
	res, err = p.Run(&evm, mockEVMAddr, append(p.ExecuteID, args...))
	require.NotNil(t, err)
	args, _ = executeMethod.Inputs.Pack(contractAddr.String(), []byte("{\"echo\":{\"message\":\"test msg\"}}"), []byte("bad coins"))
	res, err = p.Run(&evm, mockEVMAddr, append(p.ExecuteID, args...))
	require.NotNil(t, err)
}

func TestQuery(t *testing.T) {
	testApp := app.Setup(false, false)
	mockAddr, _ := testkeeper.MockAddressPair()
	ctx := testApp.GetContextForDeliverTx([]byte{}).WithBlockTime(time.Now())
	wasmKeeper := wasmkeeper.NewDefaultPermissionKeeper(testApp.WasmKeeper)
	p, err := wasmd.NewPrecompile(&testApp.EvmKeeper, wasmKeeper, testApp.WasmKeeper)
	require.Nil(t, err)
	code, err := os.ReadFile("../../example/cosmwasm/echo/artifacts/echo.wasm")
	require.Nil(t, err)
	codeID, err := wasmKeeper.Create(ctx, mockAddr, code, nil)
	require.Nil(t, err)
	contractAddr, _, err := wasmKeeper.Instantiate(ctx, codeID, mockAddr, mockAddr, []byte("{}"), "test", sdk.NewCoins())
	require.Nil(t, err)

	queryMethod, err := p.ABI.MethodById(p.QueryID)
	require.Nil(t, err)
	args, err := queryMethod.Inputs.Pack(contractAddr.String(), []byte("{\"info\":{}}"))
	require.Nil(t, err)
	statedb := state.NewDBImpl(ctx, &testApp.EvmKeeper, true)
	evm := vm.EVM{
		StateDB: statedb,
	}
	res, err := p.Run(&evm, common.Address{}, append(p.QueryID, args...))
	require.Nil(t, err)
	outputs, err := queryMethod.Outputs.Unpack(res)
	require.Nil(t, err)
	require.Equal(t, 1, len(outputs))
	require.Equal(t, "{\"message\":\"query test\"}", string(outputs[0].([]byte)))

	// bad contract address
	args, _ = queryMethod.Inputs.Pack(mockAddr.String(), []byte("{\"info\":{}}"))
	res, err = p.Run(&evm, common.Address{}, append(p.ExecuteID, args...))
	require.NotNil(t, err)

	// bad input
	args, _ = queryMethod.Inputs.Pack("not bech32", []byte("{\"info\":{}}"))
	res, err = p.Run(&evm, common.Address{}, append(p.ExecuteID, args...))
	require.NotNil(t, err)
	args, _ = queryMethod.Inputs.Pack(contractAddr.String(), []byte("{\"bad\":{}}"))
	res, err = p.Run(&evm, common.Address{}, append(p.ExecuteID, args...))
	require.NotNil(t, err)
}
