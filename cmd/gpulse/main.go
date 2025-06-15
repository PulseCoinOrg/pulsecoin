package main

import (
	"time"

	"github.com/PulseCoinOrg/pulsecoin/console"
	"github.com/PulseCoinOrg/pulsecoin/core"
	"github.com/PulseCoinOrg/pulsecoin/core/types"
	"github.com/charmbracelet/log"
)

func main() {
	chain, err := core.NewBlockChain()
	if err != nil {
		log.Error(err.Error())
	}

	// insert genesis block here for now...
	err = chain.InsertOne(types.NewBlock(
		time.Now().UnixNano(),
		[]*types.Transaction{},
	))
	if err != nil {
		log.Error(err.Error())
	}

	console := console.New(chain)
	if err := console.Run(); err != nil {
		log.Error(err.Error())
	}
}
