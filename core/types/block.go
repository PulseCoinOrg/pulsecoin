package types

import (
	"bytes"
	"encoding/gob"

	"github.com/PulseCoinOrg/pulsecoin/common"
)

type Block struct {
	Time         int64
	Hash         common.Hash
	ParentHash   common.Hash
	Transactions []string // TODO change this to []*Transaction
}

func NewBlock(time int64, parentHash common.Hash, txs []string) *Block {
	block := &Block{
		Time:         time,
		ParentHash:   parentHash,
		Transactions: txs,
	}
	block.Hash = common.Sha256Hash(block.Bytes())
	return block
}

// converts a block into bytes
func (b *Block) Bytes() []byte {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	if err := enc.Encode(b); err != nil {
		panic(err)
	}
	return buf.Bytes()
}

// decodes the block from bytes
func DecodeBlock(data []byte) *Block {
	var block Block
	dec := gob.NewDecoder(bytes.NewReader(data))
	if err := dec.Decode(&block); err != nil {
		panic(err)
	}
	return &block
}
