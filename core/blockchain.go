package core

import (
	"bytes"
	"fmt"
	"time"

	"github.com/PulseCoinOrg/pulsecoin/common"
	"github.com/PulseCoinOrg/pulsecoin/core/types"
	"github.com/PulseCoinOrg/pulsecoin/pulsedb/leveldb"
)

type BlockChain struct {
	Database    *leveldb.Database
	writeBuffer map[common.Hash]*types.Block
	lastBlock   *types.Block
}

func NewBlockChain() (*BlockChain, error) {
	db, err := leveldb.New("chaindb-out")
	if err != nil {
		return nil, err
	}
	return &BlockChain{
		Database:    db,
		writeBuffer: make(map[common.Hash]*types.Block),
	}, nil
}

func createGenesis() *types.Block {
	genesis := types.NewBlock(time.Now().UnixNano(), []*types.Transaction{})
	return genesis
}

func (chain *BlockChain) InsertOne(block *types.Block) error {
	if chain.Database == nil {
		return ErrChainDatabaseClosed
	}
	if chain.writeBuffer == nil {
		return ErrChainWriteBufferClosed
	}

	// Genesis case
	if chain.lastBlock == nil {
		genesis := createGenesis()
		fmt.Printf("Genesis block created at %d (UNIX), Hash: %s\n", time.Now().Unix(), genesis.Hash.String())
		chain.writeBuffer[genesis.Hash] = genesis
		if err := chain.Database.Put(genesis.Hash.Bytes(), genesis.Bytes()); err != nil {
			return err
		}
		chain.lastBlock = genesis
		block.ParentHash = genesis.Hash
	} else {
		block.ParentHash = chain.lastBlock.Hash
	}

	chain.writeBuffer[block.Hash] = block
	if err := chain.Database.Put(block.Hash.Bytes(), block.Bytes()); err != nil {
		return err
	}

	// Update latest block
	chain.lastBlock = block

	return nil
}

// inserts a chain of blocks into the blockchain
func (chain *BlockChain) InsertMany(blocks []*types.Block) error {
	if chain.Database == nil {
		return ErrChainDatabaseClosed
	}
	if chain.writeBuffer == nil {
		return ErrChainWriteBufferClosed
	}

	for _, block := range blocks {
		// manually trigger DB flush to update latest block before inserting the next
		if err := chain.InsertOne(block); err != nil {
			return err
		}
	}

	return nil
}

// verification that all blocks are valid
func (chain *BlockChain) SanityCheck() bool {
	if chain.Database == nil {
		return false
	}

	iter := chain.Database.NewIterator(nil)
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

func (chain *BlockChain) BlockByHash(hash string) *types.Block {
	for h, block := range chain.writeBuffer {
		if h.String() == hash {
			return block
		}
	}

	return nil
}

func (chain *BlockChain) Copy() *BlockChain {
	newBuffer := make(map[common.Hash]*types.Block)
	for hash, block := range chain.writeBuffer {
		newBlock := *block  // copy the value
		copied := &newBlock // get pointer to copied block
		newBuffer[hash] = copied
	}

	var newLastBlock *types.Block
	if chain.lastBlock != nil {
		lastCopy := *chain.lastBlock
		newLastBlock = &lastCopy
	}

	return &BlockChain{
		Database:    chain.Database,
		writeBuffer: newBuffer,
		lastBlock:   newLastBlock,
	}
}
