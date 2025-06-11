package types

import (
	"github.com/PulseCoinOrg/pulsecoin/common"
)

type Transaction struct {
	FromAddr      common.Address
	RecipientAddr common.Address
	Charge        int64
	Time          int64
}

func NewTx(from common.Address, to common.Address, charge int64, time int64) *Transaction {
	return &Transaction{
		FromAddr:      from,
		RecipientAddr: to,
		Charge:        charge,
		Time:          time,
	}
}
