package db

import (
	"bytes"

	"github.com/dgraph-io/badger"
	"github.com/rescale-labs/scaleshift/api/src/config"
)

var (
	db *badger.DB
)

func init() {
	opts := badger.DefaultOptions
	opts.Dir = config.Config.DatabaseDir
	opts.ValueDir = config.Config.DatabaseDir
	if candidate, err := badger.Open(opts); err == nil {
		db = candidate
	}
	if db == nil {
		panic("DB cannot be created")
	}
}

// ShutdownDB stop its service
func ShutdownDB() {
	db.Close()
}

// SetValue with the specified function
func SetValue(f func(txn *badger.Txn) error) error {
	return db.Update(f)
}

// GetValue with the specified function
func GetValue(f func(txn *badger.Txn) error) error {
	return db.View(f)
}

// GetValueSimple returns []bytes
func GetValueSimple(key string) ([]byte, error) {
	out := bytes.Buffer{}
	if err := GetValue(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(key))
		if err != nil {
			if err == badger.ErrKeyNotFound {
				return nil // ignore errors
			}
			return err
		}
		value, err := item.Value()
		if err != nil {
			return err
		}
		out.Write(value)
		return nil
	}); err != nil {
		return nil, err
	}
	return out.Bytes(), nil
}
