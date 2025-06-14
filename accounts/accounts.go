package accounts

import (
	"errors"

	"github.com/PulseCoinOrg/pulsecoin/common"
	"github.com/PulseCoinOrg/pulsecoin/pulsedb/memorydb"
)

var (
	ErrWriteBufferClosed = errors.New("accounts write buffer is closed")
)

type Accounts struct {
	memory      *memorydb.Database
	writeBuffer map[common.Address]*Wallet
}

func NewManager() *Accounts {
	return &Accounts{
		memory:      memorydb.New(),
		writeBuffer: make(map[common.Address]*Wallet),
	}
}

func (acs *Accounts) InsertWallet(wallet *Wallet) error {
	if acs.writeBuffer == nil {
		return ErrWriteBufferClosed
	}
	acs.writeBuffer[wallet.Address] = wallet

	w := acs.writeBuffer[wallet.Address]

	err := acs.memory.Put(w.Address.Bytes(), w.Bytes())
	if err != nil {
		return err
	}

	return nil
}
