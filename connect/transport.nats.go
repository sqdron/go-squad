package connect

import (
	"github.com/nats-io/nats"
	"encoding/json"
	"time"
	"fmt"
	"reflect"
	"log"
	"github.com/sqdron/squad/util"
)

type transport struct {
	connection *nats.Conn
}

func (t *transport) Subscribe(s string, cb interface{}) {
	t.QueueSubscribe(s, "", cb)
}

func (t *transport) Publish(s string, message interface{}) error{
	data, e := json.Marshal(message)
	if (e != nil){
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
	data, _ := json.Marshal(message)
	msg, e := t.connection.Request(s, data, timout)
	fmt.Println(msg)
	if (e != nil) {
		return nil, e
	}
	return msg.Data, nil
}

func NatsTransport(url string) ITransport {
	nc, e := nats.Connect(url)
	if e != nil {
		panic(e)
	}
	log.Println("Nats client started at: " + url)

	return &transport{connection: nc}
}

func (t *transport) QueueSubscribe(s string, group string, cb interface{}) {

	dispatcher := &messageDispatcher{}
	dispatcher.Use(UnmarshalMessageFor(cb))
	dispatcher.Use(ApplyCallback(cb))

	t.connection.QueueSubscribe(s, group, func(m *nats.Msg) {
		dispatcher.Apply(m)
		if m.Reply != "" {
			var data []byte
			switch responce.(type) {
			default:
				data, _ = json.Marshal(responce)
			case []byte:
				data = responce.([]byte)
			}
			fmt.Println(data)
			t.connection.Publish(m.Reply, data)
		}
	})
}
