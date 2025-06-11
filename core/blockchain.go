package core

import "github.com/PulseCoinOrg/pulsecoin/pulsedb/leveldb"

type BlockChain struct {
	db *leveldb.Database
}
