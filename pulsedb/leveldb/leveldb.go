package leveldb

import "github.com/syndtr/goleveldb/leveldb"

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
