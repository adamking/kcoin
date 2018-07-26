package konsensus

import (
	"math/big"

	"github.com/kowala-tech/kcoin/client/common"
)

type OracleMgr interface {
	AveragePrice() (*big.Int, error)
	Submissions() ([]common.Address, error)
}

type System interface {
	CurrencyPrice() (*big.Int, error)
	CurrencySupply() (*big.Int, error)
	MintedAmount() (*big.Int, error)
	OracleDeduction(*big.Int) (*big.Int, error)
	OracleReward(*big.Int) (*big.Int, error)
	Address() common.Address
}