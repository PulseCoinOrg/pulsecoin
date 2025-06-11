package main

import (
	"log/slog"
	"time"

	"github.com/PulseCoinOrg/pulsecoin/common"
	"github.com/PulseCoinOrg/pulsecoin/core"
	"github.com/PulseCoinOrg/pulsecoin/core/types"
)

func main() {
	chain, err := core.NewBlockChain()
	if err != nil {
		slog.Error("failed to create chain", "err", err.Error())
	}
	if chain == nil {
		slog.Error("failed to create chain", "err", err.Error())
	}
	hash := common.Sha256Hash([]byte("Hello World!"))
	block0 := types.NewBlock(time.Now().UnixNano(), hash, []string{})
	block1 := types.NewBlock(time.Now().UnixNano(), block0.Hash, []string{})
	block2 := types.NewBlock(time.Now().UnixNano(), block1.Hash, []string{})
	block3 := types.NewBlock(time.Now().UnixNano(), block2.Hash, []string{})

	err = chain.InsertMany([]*types.Block{block0, block1, block2, block3})
	if err != nil {
		slog.Error("failed to insert many blocks", "err", err.Error())
	}

	for _, block := range []*types.Block{block0, block1, block2, block3} {
		slog.Info("block has been inserted", "hash", block.Hash.String())
	}

	isSane := chain.SanityCheck()
	if isSane == false {
		slog.Error("chain is invalid")
	}
}
