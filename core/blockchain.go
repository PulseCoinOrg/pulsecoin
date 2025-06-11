package core

import (
	"github.com/PulseCoinOrg/pulsecoin/common"
	"github.com/PulseCoinOrg/pulsecoin/core/types"
	"github.com/PulseCoinOrg/pulsecoin/pulsedb/leveldb"
)

type BlockChain struct {
	db          *leveldb.Database
	writeBuffer map[common.Hash]*types.Block
}

func NewBlockChain() (*BlockChain, error) {
	db, err := leveldb.New("chain")
	if err != nil {
		return nil, err
	}
	return &BlockChain{
		db:          db,
		writeBuffer: make(map[common.Hash]*types.Block),
	}, nil
}

func (chain *BlockChain) InsertBlock(block *types.Block) error {
	if chain.db == nil {
		return ErrChainDatabaseClosed
	}

	if chain.writeBuffer == nil {
		return ErrChainWriteBufferClosed
	}

	chain.writeBuffer[block.Hash] = block

	// TODO change this line to be more clean
	// key := ...
	// value := ...
	err := chain.db.Put(chain.writeBuffer[block.Hash].Hash.Bytes(), chain.writeBuffer[block.Hash].Bytes())
	if err != nil {
		return err
	}

	return nil
}

func (chain *BlockChain) SanityCheck() bool {
	return true
}
