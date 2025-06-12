package leveldb

import (
	"errors"

	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/iterator"
	"github.com/syndtr/goleveldb/leveldb/util"
)

var (
	ErrNotFound = errors.New("not found in key-value store")
)

type Database struct {
	db *leveldb.DB
}

// New() creates a new instance of leveldb and returns a wrapper struct around leveldb
func New(file string) (*Database, error) {
	db, err := leveldb.OpenFile(file, nil)
	if err != nil {
		return nil, err
	}
	return &Database{
		db: db,
	}, nil
}

func (db *Database) Get(key []byte) ([]byte, error) {
	return db.db.Get(key, nil)
}

func (db *Database) Put(key, value []byte) error {
	return db.db.Put(key, value, nil)
}

func (db *Database) Delete(key []byte) error {
	return db.db.Delete(key, nil)
}

// returns true if the key-value store has a specific key
func (db *Database) Has(key []byte) (bool, error) {
	return db.db.Has(key, nil)
}

// gets the previous item's value in the key-value store
func (db *Database) PreviousItem() ([]byte, error) {
	iter := db.NewIterator(nil)
	defer iter.Release()

	if !iter.Last() {
		return nil, ErrNotFound
	}

	return iter.Value(), nil
}

type Iterator struct {
	iter iterator.Iterator
}

func (db *Database) NewIterator(prefix []byte) *Iterator {
	r := util.BytesPrefix(prefix)
	return &Iterator{
		iter: db.db.NewIterator(r, nil),
	}
}

// returns true based on the first element inserted
func (it *Iterator) First() bool {
	return it.iter.First()
}

// returns true if the last element is found
func (it *Iterator) Last() bool {
	return it.iter.Last()
}

// returns true if the previous item is found
func (it *Iterator) Prev() bool {
	return it.iter.Prev()
}

// returns true if there is a next item in memory
func (it *Iterator) Next() bool {
	return it.iter.Next()
}

// gets the key from the iterator
func (it *Iterator) Key() []byte {
	return it.iter.Key()
}

// gets the value from the iterator
func (it *Iterator) Value() []byte {
	return it.iter.Value()
}

// returns true if the iterator is valid
func (it *Iterator) Valid() bool {
	return it.iter.Valid()
}

// retrieves an iterator's error if any
func (it *Iterator) Error() error {
	return it.iter.Error()
}

func (it *Iterator) Release() {
	it.iter.Release()
}
