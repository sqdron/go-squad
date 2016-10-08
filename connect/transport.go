package connect

import (
	"github.com/nats-io/nats"
	"github.com/sqdron/squad/middleware"
	"reflect"
	"github.com/sqdron/squad/util"
	"encoding/json"
	"fmt"
	"log"
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
	var res = reflect.ValueOf(ctx.action).Call(content)[0].Interface()
	return res
}

type ITransport interface {
	Subscribe(s string, cb interface{})
	QueueSubscribe(s string, group string, cb interface{})
	Request(s string, message interface{}, cb interface{}) error
}

func NewTransport(url string) ITransport {

	nc, e := nats.Connect(url)
	if (e != nil){
		panic(e)
	}
	log.Println("Application started at: " + url)
	return &transport{connection: nc}
}

func (t *transport) Subscribe(s string, cb interface{}) {
	t.QueueSubscribe(s, "", cb)
}

func (t *transport) QueueSubscribe(s string, group string, cb interface{}) {
	fmt.Println("Subscrive for " + s)
	requestType := reflect.TypeOf(cb).In(0);
	requestObject := reflect.New(requestType).Elem()

	h := middleware.ApplyMiddleware(CreateEncoderMiddleware(requestObject.Interface()))

	t.connection.QueueSubscribe(s, group, func(m *nats.Msg) {
		fmt.Printf("Handling message (%s) and reply to %s\n", m.Subject, m.Reply)
		ctx := &requestContext{action: cb}
		h(ctx).Apply(m)
	})
}

func (t *transport) Request(s string, message interface{}, cb interface{}) error {
	fmt.Print("Submit activation data: ")
	fmt.Println(message)
	data, _ := json.Marshal(message)
	replay := "reply_" + util.GenerateString(10)
	t.Subscribe(replay, cb)
	fmt.Printf("Publish to (%s) with reply %s \n", s, replay)
	return t.connection.PublishRequest(s, replay, data)
}
