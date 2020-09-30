package db

import (
	"bytes"
	"time"

	"github.com/dgraph-io/badger"
	"github.com/scaleshift/scaleshift/api/src/config"
)

var (
	cache *badger.DB
)

func init() {
	opts := badger.DefaultOptions(config.Config.DatabaseDir)
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
		entry := &badger.Entry{
			Key:   []byte(key),
			Value: value,
		}
		if duration != nil {
			entry.WithTTL(*duration)
		}
		return txn.SetEntry(entry)
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
		var value []byte
		if err = item.Value(func(val []byte) error {
			value = append([]byte{}, val...)
			return nil
		}); err != nil {
			return err
		}
		out.Write(value)
		return nil
	}); err != nil {
		return nil, err
	}
	return out.Bytes(), nil
}
