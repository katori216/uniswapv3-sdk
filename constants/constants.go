package constants

import (
	"math/big"

	"github.com/daoleno/uniswap-sdk-core/entities"
	"github.com/ethereum/go-ethereum/common"
)

const PoolInitCodeHash = "0xe34f199b19b2b4f47f68442619d555527d244f78a3297ea89325f843f87b8b54"

var (
	FactoryAddress = common.HexToAddress("0x1F98431c8aD98523631AE4a59f267346ea31F984")
	AddressZero    = common.HexToAddress("0x0000000000000000000000000000000000000000")
)

// The default factory enabled fee amounts, denominated in hundredths of bips.
type FeeAmount uint64

const (
	F100    FeeAmount = 100
	F400    FeeAmount = 400
	F500    FeeAmount = 500
	F1500   FeeAmount = 1500
	F2500   FeeAmount = 2500
	F3000   FeeAmount = 3000
	F10000  FeeAmount = 10000
	F50000  FeeAmount = 50000
	F100000 FeeAmount = 1000000
)

// The default factory tick spacings by fee amount.
var TickSpacings = map[FeeAmount]int{
	F100:    1,
	F400:    8,
	F500:    10,
	F1500:   30,
	F2500:   50,
	F3000:   60,
	F10000:  200,
	F50000:  1000,
	F100000: 2000,
}

var (
	NegativeOne = big.NewInt(-1)
	Zero        = big.NewInt(0)
	One         = big.NewInt(1)

	// used in liquidity amount math
	Q96  = new(big.Int).Exp(big.NewInt(2), big.NewInt(96), nil)
	Q192 = new(big.Int).Exp(Q96, big.NewInt(2), nil)

	PercentZero = entities.NewFraction(big.NewInt(0), big.NewInt(1))
)
