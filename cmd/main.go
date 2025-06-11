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
	block := types.NewBlock(time.Now().UnixNano(), hash, []string{})

	err = chain.InsertOne(block)
	if err != nil {
		slog.Error("failed to insert one block into the chain", "hash", block.Hash.String())
	}

	slog.Info("inserted one block into the chain...", "hash", block.Hash.String())
}
