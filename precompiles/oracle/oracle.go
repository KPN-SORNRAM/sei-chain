package oracle

import (
	"bytes"
	"embed"
	"math/big"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
	pcommon "github.com/sei-protocol/sei-chain/precompiles/common"
	"github.com/sei-protocol/sei-chain/x/oracle/types"
)

const (
	GetExchangeRatesMethod = "getExchangeRates"
	GetOracleTwapsMethod   = "getOracleTwaps"
)

const (
	OracleAddress = "0x0000000000000000000000000000000000001008"
)

var _ vm.PrecompiledContract = &Precompile{}

// Embed abi json file to the executable binary. Needed when importing as dependency.
//
//go:embed abi.json
var f embed.FS

func GetABI() abi.ABI {
	abiBz, err := f.ReadFile("abi.json")
	if err != nil {
		panic(err)
	}

	newAbi, err := abi.JSON(bytes.NewReader(abiBz))
	if err != nil {
		panic(err)
	}
	return newAbi
}

type Precompile struct {
	pcommon.Precompile
	evmKeeper    pcommon.EVMKeeper
	oracleKeeper pcommon.OracleKeeper
	address      common.Address

	GetExchangeRatesId []byte
	GetOracleTwapsId   []byte
}

// Define types which deviate slightly from cosmos types (ExchangeRate string vs sdk.Dec)
type OracleExchangeRate struct {
	ExchangeRate        string `json:"exchangeRate"`
	LastUpdate          string `json:"lastUpdate"`
	LastUpdateTimestamp int64  `json:"lastUpdateTimestamp"`
}

type DenomOracleExchangeRatePair struct {
	Denom                 string             `json:"denom"`
	OracleExchangeRateVal OracleExchangeRate `json:"oracleExchangeRateVal"`
}

type OracleTwap struct {
	Denom           string `json:"denom"`
	Twap            string `json:"twap"`
	LookbackSeconds int64  `json:"lookbackSeconds"`
}

func NewPrecompile(oracleKeeper pcommon.OracleKeeper, evmKeeper pcommon.EVMKeeper) (*Precompile, error) {
	newAbi := GetABI()

	p := &Precompile{
		Precompile:   pcommon.Precompile{ABI: newAbi},
		evmKeeper:    evmKeeper,
		address:      common.HexToAddress(OracleAddress),
		oracleKeeper: oracleKeeper,
	}

	for name, m := range newAbi.Methods {
		switch name {
		case GetExchangeRatesMethod:
			p.GetExchangeRatesId = m.ID
		case GetOracleTwapsMethod:
			p.GetOracleTwapsId = m.ID
		}
	}

	return p, nil
}

// RequiredGas returns the required bare minimum gas to execute the precompile.
func (p Precompile) RequiredGas(input []byte) uint64 {
	methodID := input[:4]

	method, err := p.ABI.MethodById(methodID)
	if err != nil {
		// This should never happen since this method is going to fail during Run
		return 0
	}

	return p.Precompile.RequiredGas(input, p.IsTransaction(method.Name))
}

func (p Precompile) Address() common.Address {
	return p.address
}

func (p Precompile) Run(evm *vm.EVM, _ common.Address, input []byte, value *big.Int) (bz []byte, err error) {
	ctx, method, args, err := p.Prepare(evm, input)
	if err != nil {
		return nil, err
	}

	switch method.Name {
	case GetExchangeRatesMethod:
		p.getExchangeRates(ctx, method, args, value)
		return p.getExchangeRates(ctx, method, args, value)
	case GetOracleTwapsMethod:
		return p.getOracleTwaps(ctx, method, args, value)
	}
	return
}

func (p Precompile) getExchangeRates(ctx sdk.Context, method *abi.Method, args []interface{}, value *big.Int) ([]byte, error) {
	pcommon.AssertNonPayable(value)
	exchangeRates := []DenomOracleExchangeRatePair{}
	p.oracleKeeper.IterateBaseExchangeRates(ctx, func(denom string, rate types.OracleExchangeRate) (stop bool) {
		exchangeRates = append(exchangeRates, DenomOracleExchangeRatePair{Denom: denom, OracleExchangeRateVal: OracleExchangeRate{ExchangeRate: rate.ExchangeRate.String(), LastUpdate: rate.LastUpdate.String(), LastUpdateTimestamp: rate.LastUpdateTimestamp}})
		return false
	})

	return method.Outputs.Pack(exchangeRates)
}

func (p Precompile) getOracleTwaps(ctx sdk.Context, method *abi.Method, args []interface{}, value *big.Int) ([]byte, error) {
	pcommon.AssertNonPayable(value)
	pcommon.AssertArgsLength(args, 1)
	lookbackSeconds := args[0].(uint64)
	twaps, err := p.oracleKeeper.CalculateTwaps(ctx, lookbackSeconds)
	if err != nil {
		return nil, err
	}
	// Convert twap to string
	var oracleTwaps []OracleTwap
	for _, twap := range twaps {
		oracleTwaps = append(oracleTwaps, OracleTwap{Denom: twap.Denom, Twap: twap.Twap.String(), LookbackSeconds: twap.LookbackSeconds})
	}
	return method.Outputs.Pack(oracleTwaps)
}

func (Precompile) IsTransaction(string) bool {
	return false
}
