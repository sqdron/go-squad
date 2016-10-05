package nats

import (
	"github.com/sqdron/squad/endpoint"
	"github.com/nats-io/nats"
	"time"
	"encoding/json"
	"github.com/sqdron/squad/util"
)

type natsTransport struct {
	connection *nats.EncodedConn
}

func NatsEndpoint(url string) *endpoint.Endpoint {
	nc, _ := nats.Connect(url)
	ec, _ := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	return endpoint.NewEndpoint(&natsTransport{connection:ec})
}

//TODO: refactor this
func (t *natsTransport) Request(subject string, data interface{}) <- chan interface{} {
	result := make(chan interface{})
	sender := make(chan *endpoint.Message)
	t.connection.BindSendChan(subject, sender)

	message := &endpoint.Message{
		ID:util.GenerateString(10),
		Responce: "resp_" + util.GenerateString(5)}

	d, _ := json.Marshal(data)
	message.Payload = string(d)
	go func() {
		receiver := make(chan *endpoint.Message)
		t.connection.BindRecvChan(message.Responce, receiver)
		select {
		case response := <-receiver:
			var respData interface{}
			err := json.Unmarshal([]byte(response.Payload), &respData)
			if err != nil {
				panic(err)
			}

			result <-respData
		case <-time.After(3 * time.Second):
		//TODO: Use better way for handdling errors
			panic("Request timeout error")
		}
	}()

	sender <- message
	return result
}

//TODO: refactor this
func (t *natsTransport) Listen(subject string, handler interface{}) {
	receiver := make(chan *endpoint.Message)
	t.connection.BindRecvChan(subject, receiver)
	go func(){
		for{
			message := <- receiver
			var data interface{}
			err := json.Unmarshal([]byte(message.Payload), &data)
			if err != nil {
				panic(err)
			}
			result := message.Apply(handler, data)
			responceMessage := &endpoint.Message{
				ID:util.GenerateString(10)}
			//

			d, _ := json.Marshal(result)
			responceMessage.Payload = string(d)
			sender := make(chan *endpoint.Message)
			t.connection.BindSendChan(message.Responce, sender)
			sender <- responceMessage
		}
	}()
}

func (t *natsTransport) Close() {
	t.connection.Close()
}
