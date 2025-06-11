package memorydb

import (
	"errors"
	"sync"
)

var (
	errMemoryDbClosed = errors.New("memorydb is closed")
)

type Database struct {
	db   map[string][]byte
	lock sync.RWMutex
}

func New() *Database {
	return &Database{
		db: make(map[string][]byte),
	}
}

// creates a new key-value store with a size cap
func NewWithLimit(limit int) *Database {
	return &Database{
		db: make(map[string][]byte, limit),
	}
}

// closes the in-memory key-value store
func (db *Database) Close() error {
	db.lock.Lock()
	defer db.lock.Unlock()
	db.db = nil
	return nil
}

// retrieves an item from the in-memory key-value store
func (db *Database) Get(key []byte) ([]byte, error) {
	db.lock.RLock()
	defer db.lock.RUnlock()
	if db.db == nil {
		return nil, errMemoryDbClosed
	}
	return db.db[string(key)], nil
}

// inserts an item into the in-memory key-value store
func (db *Database) Put(key, value []byte) error {
	db.lock.Lock()
	defer db.lock.Unlock()
	if db.db == nil {
		return errMemoryDbClosed
	}
	db.db[string(key)] = value
	return nil
}

// deletes an item from the in-memory key-value store
func (db *Database) Delete(key []byte) error {
	db.lock.Lock()
	defer db.lock.Unlock()
	if db.db == nil {
		return errMemoryDbClosed
	}
	delete(db.db, string(key))
	return nil
}
