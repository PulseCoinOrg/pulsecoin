package core

import (
	"bytes"

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

// inserts just one block into the blockchain
func (chain *BlockChain) InsertOne(block *types.Block) error {
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
	err := chain.db.Put(
		chain.writeBuffer[block.Hash].Hash.Bytes(),
		chain.writeBuffer[block.Hash].Bytes(),
	)
	if err != nil {
		return err
	}

	return nil
}

// inserts a chain of blocks into the blockchain
func (chain *BlockChain) InsertMany(blocks []*types.Block) error {
	if chain.db == nil {
		return ErrChainDatabaseClosed
	}
	if chain.writeBuffer == nil {
		return ErrChainWriteBufferClosed
	}
	var err error
	for _, block := range blocks {
		err = chain.InsertOne(block)
	}

	if err != nil {
		return err
	}

	return nil
}

// verification that all blocks are valid
func (chain *BlockChain) SanityCheck() bool {
	if chain.db == nil {
		return false
	}

	iter := chain.db.NewIterator(nil)
	defer iter.Release()

	seen := make(map[common.Hash]bool)

	for iter.First(); iter.Valid(); iter.Next() {
		key := iter.Key()
		value := iter.Value()

		block := types.DecodeBlock(value)

		if !bytes.Equal(key, block.Hash.Bytes()) {
			return false
		}

		if block.ParentHash != chain.writeBuffer[block.Hash].ParentHash {
			return false
		}

		seen[block.Hash] = true
	}

	if err := iter.Error(); err != nil {
		return false
	}

	return true
}
