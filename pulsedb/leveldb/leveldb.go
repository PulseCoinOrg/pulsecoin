package leveldb

import (
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/iterator"
	"github.com/syndtr/goleveldb/leveldb/util"
)

type Database struct {
	db *leveldb.DB
}

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

func (db *Database) Has(key []byte) (bool, error) {
	return db.db.Has(key, nil)
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

func (it *Iterator) First() bool {
	return it.iter.First()
}

func (it *Iterator) Next() bool {
	return it.iter.Next()
}

func (it *Iterator) Key() []byte {
	return it.iter.Key()
}

func (it *Iterator) Value() []byte {
	return it.iter.Value()
}

func (it *Iterator) Valid() bool {
	return it.iter.Valid()
}

func (it *Iterator) Error() error {
	return it.iter.Error()
}

func (it *Iterator) Release() {
	it.iter.Release()
}
