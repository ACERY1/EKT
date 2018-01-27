package db

import (
	_ "github.com/mattn/go-sqlite3"
	_ "github.com/syndtr/goleveldb/leveldb"
)

type EKTDB interface {
	Set(key, value []byte)
	Get(Key, value []byte)
	Delete(Key []byte)
}

func init() {
	//leveldb.Open()
}
