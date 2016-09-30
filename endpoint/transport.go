package endpoint

import "context"

type transport struct {
	provider interface{}
	listen   Middleware
}

type ITransport interface {
	Configuratoin() []interface{}
	//Publish() chan <- Message
	Listen() <-chan Message
}

type MessageHandler func(ctx context.Context, request interface{}) (response Message, err error)
type Middleware func(MessageHandler) MessageHandler

func NewTransport(provider interface{}) *transport {
	return &transport{provider: provider}
}

func (t *transport) Configuratoin() []interface{} {
	return t.provider.(ITransport).Configuratoin()
}

func (t *transport) Listen() <-chan Message {
	out := make(chan Message)
	go func() {
		defer close(out)
		for {
			select {
			case m := <-t.provider.(ITransport).Listen():
				out <- t.listen(nil, m)
			}
		}
		t.listen()
	}()
	return t.provider.(ITransport).Listen()
}
