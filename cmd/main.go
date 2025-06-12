package main

import (
	"fmt"
	"time"

	"github.com/PulseCoinOrg/pulsecoin/core"
	"github.com/PulseCoinOrg/pulsecoin/core/types"
	"github.com/charmbracelet/log"
)

func main() {
	chain, err := core.NewBlockChain()
	if err != nil {
		log.Error(err.Error())
		return
	}

	block1 := types.NewBlock(time.Now().UnixNano(), []*types.Transaction{})
	block2 := types.NewBlock(time.Now().UnixNano(), []*types.Transaction{})
	block3 := types.NewBlock(time.Now().UnixNano(), []*types.Transaction{})

	err = chain.InsertMany([]*types.Block{block1, block2, block3})
	if err != nil {
		log.Error(err.Error())
		return
	}

	fmt.Println("block #1: ", block1.MarshalJSON())
	fmt.Println("block #2: ", block2.MarshalJSON())
	fmt.Println("block #3: ", block3.MarshalJSON())
}
