package db

import (
	"time"

	"github.com/go-msgqueue/msgqueue"
	"github.com/go-msgqueue/msgqueue/memqueue"
	"github.com/xtream-d-labs/ai-gateway/api/src/config"
)

var q *memqueue.Queue

// SetupQueue start its service
func SetupQueue(handler interface{}, fallbackHandler interface{}) {
	q = memqueue.NewQueue(&msgqueue.Options{
		Handler:         handler,
		FallbackHandler: fallbackHandler,
		MaxWorkers:      config.Config.MaxWorkers,
		MaxFetchers:     config.Config.MaxFetchers,
		RetryLimit:      1,
	})
}

// Enqueue set a job in the queue
func Enqueue(message string) error {
	return q.Add(msgqueue.NewMessage(message))
}

// ShutdownQueue stop its service
func ShutdownQueue() {
	q.CloseTimeout(5 * time.Second)
}
