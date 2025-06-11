package main

import (
	"fmt"

	"github.com/PulseCoinOrg/pulsecoin/pulsedb/leveldb"
)

func main() {
	db, err := leveldb.New("chain")
	if err != nil {
		panic(err)
	}

	err = db.Put([]byte("1"), []byte("hello world"))
	err = db.Put([]byte("2"), []byte("world hello"))

	iter := db.NewIterator([]byte(""))
	defer iter.Release()

	for ok := iter.First(); ok; ok = iter.Next() {
		fmt.Printf("key = %s, value = %s\n", iter.Key(), iter.Value())
	}
}
