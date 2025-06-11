package main

import (
	"fmt"
	"time"

	"github.com/PulseCoinOrg/pulsecoin/common"
	"github.com/PulseCoinOrg/pulsecoin/core/types"
)

func main() {
	parentHash := common.Sha256Hash([]byte("random-data"))
	block := types.NewBlock(time.Now().UnixNano(), parentHash, []string{})
	fmt.Println(block.Hash.String())
}
