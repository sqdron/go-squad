package connect

import (
	"encoding/json"
	"fmt"
	"github.com/nats-io/nats"
	"github.com/sqdron/squad/util"
	"log"
	"time"
)

type transport struct {
	connection *nats.Conn
}

func NatsTransport(url string) ITransport {
	nc, e := nats.Connect(url)
	if e != nil {
		panic(e)
	}
	return &transport{connection: nc}
}

func (t *transport) Subscribe(s string, cb interface{}) {
	t.QueueSubscribe(s, "", cb)
}

func (t *transport) Publish(s string, message interface{}) error {
	data, e := json.Marshal(message)
	if e != nil {
		return e
	}
	return t.connection.Publish(s, data)
}

func (t *transport) Request(s string, message interface{}, cb interface{}) error {
	data, _ := json.Marshal(message)
	replay := "reply_" + util.GenerateString(10)
	t.Subscribe(replay, cb)
	return t.connection.PublishRequest(s, replay, data)
}

func (t *transport) RequestSync(s string, message interface{}, timout time.Duration) (interface{}, error) {
	data, e := marshalMessage(message)
	msg, e := t.connection.Request(s, data, timout)
	if e != nil {
		fmt.Println(e)
		return nil, e
	}
	return msg.Data, nil
}

func (t *transport) QueueSubscribe(s string, group string, cb interface{}) {
	t.connection.QueueSubscribe(s, group, func(m *nats.Msg) {
		result, e := applyMessage(s, m.Data, cb)
		if e != nil {
			log.Println(e)
		}
		if m.Reply != "" {
			data, _ := marshalMessage(result)
			t.connection.Publish(m.Reply, data)
		}
	})
}
