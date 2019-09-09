package db

import (
	"bytes"
	"time"

	"github.com/dgraph-io/badger"
	"github.com/rescale-labs/scaleshift/api/src/config"
)

var (
	cache *badger.DB
)

func init() {
	opts := badger.DefaultOptions
	opts.Dir = config.Config.DatabaseDir
	opts.ValueDir = config.Config.DatabaseDir
	if candidate, err := badger.Open(opts); err == nil {
		cache = candidate
	}
	if cache == nil {
		panic("DB cannot be created")
	}
}

// ShutdownCache stop its service
func ShutdownCache() {
	cache.Close()
}

// SetCache with the specified function
func SetCache(key string, value []byte, duration *time.Duration) error {
	return cache.Update(func(txn *badger.Txn) error {
		if duration != nil {
			return txn.SetWithTTL([]byte(key), value, *duration)
		}
		return txn.Set([]byte(key), value)
	})
}

// GetCache returns []bytes
func GetCache(key string) ([]byte, error) {
	out := bytes.Buffer{}
	if err := cache.View(func(txn *badger.Txn) error {
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
