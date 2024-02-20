package keeper_test

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/sei-protocol/sei-chain/testutil/keeper"
	"github.com/stretchr/testify/require"
)

func TestCode(t *testing.T) {
	k, ctx := keeper.MockEVMKeeper()
	_, addr := keeper.MockAddressPair()

	require.Equal(t, common.Hash{}, k.GetCodeHash(ctx, addr))
	require.Nil(t, k.GetCode(ctx, addr))
	require.Equal(t, 0, k.GetCodeSize(ctx, addr))

	code := []byte{1, 2, 3, 4, 5}
	k.SetCode(ctx, addr, code)
	require.Equal(t, crypto.Keccak256Hash(code), k.GetCodeHash(ctx, addr))
	require.Equal(t, code, k.GetCode(ctx, addr))
	require.Equal(t, 5, k.GetCodeSize(ctx, addr))
}

func TestNilCode(t *testing.T) {
	k, ctx := keeper.MockEVMKeeper()
	_, addr := keeper.MockAddressPair()

	k.SetCode(ctx, addr, nil)
	require.Nil(t, k.GetCode(ctx, addr))
	require.Equal(t, 0, k.GetCodeSize(ctx, addr))
	require.Equal(t, ethtypes.EmptyCodeHash, k.GetCodeHash(ctx, addr))
}
