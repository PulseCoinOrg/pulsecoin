package types

import (
	"bytes"
	"encoding/gob"

	"github.com/PulseCoinOrg/pulsecoin/common"
)

type Transaction struct {
	FromAddr      common.Address
	RecipientAddr common.Address
	Charge        int64
	Hash          common.Hash
	Time          int64
}

func NewTx(
	from common.Address,
	to common.Address,
	charge int64,
	time int64,
) *Transaction {
	tx := &Transaction{
		FromAddr:      from,
		RecipientAddr: to,
		Charge:        charge,
		Time:          time,
	}
	tx.Hash = common.Sha256Hash(tx.Bytes())
	return tx
}

func (tx *Transaction) Bytes() []byte {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	if err := enc.Encode(tx); err != nil {
		panic(err)
	}
	return buf.Bytes()
}

func DecodeTransaction(data []byte) *Transaction {
	var tx *Transaction
	dec := gob.NewDecoder(bytes.NewReader(data))
	if err := dec.Decode(tx); err != nil {
		panic(err)
	}
	return tx
}
