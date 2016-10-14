package connect

import (
	"time"
)

type ITransport interface {
	Subscribe(s string, cb interface{})
	QueueSubscribe(s string, group string, cb interface{})
	Publish(s string, message interface{}) error
	Request(s string, message interface{}, cb interface{}) error
	RequestSync(s string, message interface{}, timout time.Duration) (interface{}, error)
}
