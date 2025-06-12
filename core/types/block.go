package types

import (
	"bytes"
	"encoding/gob"
	"encoding/json"

	"github.com/PulseCoinOrg/pulsecoin/common"
)

type Block struct {
	Time         int64
	Hash         common.Hash
	ParentHash   common.Hash
	Transactions []*Transaction
}

func NewBlock(time int64, txs []*Transaction) *Block {
	block := &Block{
		Time:         time,
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

type MarshalBlock struct {
	Time       int64
	ParentHash string
	Hash       string
	TxHashes   []string
}

func (b *Block) MarshalJSON() string {
	var txHashes []string
	for _, tx := range b.Transactions {
		txHashes = append(txHashes, tx.Hash.String())
	}

	j, err := json.MarshalIndent(&MarshalBlock{
		Time:       b.Time,
		ParentHash: b.ParentHash.String(),
		Hash:       b.Hash.String(),
		TxHashes:   txHashes,
	}, "", "\t")
	if err != nil {
		panic(err)
	}
	return string(j)
}
