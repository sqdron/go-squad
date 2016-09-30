package nats

import (
	"github.com/sqdron/squad/endpoint"
	"github.com/nats-io/nats"
	"log"
)

type natsTransport struct {
	connection *nats.EncodedConn
}

func NatsEndpoint(url string) *endpoint.Transport {
	nc, _ := nats.Connect(url)
	ec, _ := nats.NewEncodedConn(nc, nats.JSON_ENCODER)

	return endpoint.NewEndpoint(&natsTransport{connection:ec})
}

func (t *natsTransport) Publish(subject string) chan <- *endpoint.Message {
	ch:= make(chan *endpoint.Message)
	t.connection.BindSendChan(subject, ch)
	return ch;
}

func (t *natsTransport) Listen(subject string) <- chan *endpoint.Message {
	ch := make(chan *endpoint.Message)
	log.Println("listening for:" + subject)
	t.connection.BindRecvChan(subject, ch)
	return ch
	//return t.provider.(IEndpoint).Listen(subject)
}

func (t *natsTransport) Close() {
	t.connection.Close()
}

