package connect

import (
	"encoding/json"
	"github.com/nats-io/nats"
	"github.com/sqdron/squad/middleware"
	"github.com/sqdron/squad/util"
	"log"
	"reflect"
	"time"
	"fmt"
)

type transport struct {
	connection *nats.Conn
}

type requestContext struct {
	action interface{}
}

func (ctx requestContext) Apply(r interface{}) interface{} {
	content := []reflect.Value{}
	content = append(content, reflect.ValueOf(r))
	return reflect.ValueOf(ctx.action).Call(content)[0].Interface()
}

type ITransport interface {
	Subscribe(s string, cb interface{})
	QueueSubscribe(s string, group string, cb interface{})
	Request(s string, message interface{}, cb interface{}) error
	RequestSync(s string, message interface{}, timout time.Duration) (interface{}, error)
}

func NewTransport(url string) ITransport {
	nc, e := nats.Connect(url)
	if e != nil {
		panic(e)
	}
	log.Println("Nats client started at: " + url)
	return &transport{connection: nc}
}

func (t *transport) Subscribe(s string, cb interface{}) {
	t.QueueSubscribe(s, "", cb)
}

func (t *transport) QueueSubscribe(s string, group string, cb interface{}) {
	requestType := reflect.TypeOf(cb).In(0)
	requestObject := reflect.New(requestType).Elem()

	h := middleware.ApplyMiddleware(CreateEncoderMiddleware(requestObject.Interface()))

	t.connection.QueueSubscribe(s, group, func(m *nats.Msg) {
		ctx := &requestContext{action: cb}
		responce := h(ctx).Apply(m)
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
