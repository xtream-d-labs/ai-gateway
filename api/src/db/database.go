// Package db defines data transfer obejects
package db

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/rescale-labs/scaleshift/api/src/config"
	logs "github.com/rescale-labs/scaleshift/api/src/log"
	log "github.com/sirupsen/logrus"
)

var (
	db *gorm.DB
)

// BeginTransactions starts a transaction
func BeginTransactions() *gorm.DB {
	db.Exec("set transaction isolation level serializable")
	return db.Begin()
}

// Rollback ロールバックします
func Rollback(tx *gorm.DB) {
	tx.Rollback()
}

// Commit コミットします
func Commit(tx *gorm.DB) {
	tx.Commit()
}

// Initialize inits its db settings
func Initialize() {
	impl, err := gorm.Open("mysql", config.Config.DatabaseEndpoint)
	if err != nil {
		logs.Error("cannot-open-db", err, nil)
	}
	if log.GetLevel() == log.DebugLevel {
		impl.SetLogger(wrappedLogger{log.StandardLogger()})
		impl.LogMode(true)
	}
	db = impl
	db.AutoMigrate(&Image{}, &Job{}, &Error{})
}

// Shutdown close its connection
func Shutdown() {
	if db != nil {
		db.Close() // nolint
	}
}

// Ping verifies a connection to the database is still alive,
// establishing a connection if necessary.
func Ping() error {
	if db == nil {
		return errors.New("database connection has gone")
	}
	return db.DB().Ping()
}

type wrappedLogger struct {
	original *log.Logger
}

func (w wrappedLogger) Print(values ...interface{}) {
	message := "▶︎ "
	for idx, value := range values {
		str := fmt.Sprintf("%+v", value)
		if strings.Contains(str, "database-gorm.go") || str == "sql" {
			continue
		}
		message += deeper(value, 0, 1)
		if idx < len(values)-1 {
			message += ", "
		}
	}
	w.original.Debugln(message)
}

func deeper(value interface{}, index, depth int) string {
	message := ""
	if depth > 5 {
		return message
	}
	v := reflect.ValueOf(value)
	if !v.IsValid() {
		return message
	}
	switch kind := v.Kind(); kind {
	case reflect.Array:
		message += "["
		for i := 0; i < v.Len(); i++ {
			message += deeper(v.Index(i), i, depth+1)
			if i < v.Len()-1 {
				message += ", "
			}
		}
		message += "]"

	case reflect.Slice:
		if v.IsNil() {
			return message
		}
		message += "["
		for i := 0; i < v.Len(); i++ {
			message += deeper(v.Index(i), i, depth+1)
			if i < v.Len()-1 {
				message += ", "
			}
		}
		message += "]"

	case reflect.Struct:
		candidate := fmt.Sprintf("%+v", value)
		if !strings.HasPrefix(candidate, "0x") {
			message += candidate
		} else {
			for i, n := 0, v.NumField(); i < n; i++ {
				message += fmt.Sprintf("%v", v.Field(i))
			}
		}

	case reflect.Map:
		if v.IsNil() {
			return message
		}
		message += "{"
		for idx, k := range v.MapKeys() {
			val := v.MapIndex(k)
			if !val.IsValid() || val.IsNil() {
				continue
			}
			message += fmt.Sprintf("%v: %v", k, val)
			if idx < len(v.MapKeys())-1 {
				message += ", "
			}
		}
		message += "}"

	default:
		message += fmt.Sprintf("%+v", value)
	}
	return message
}
