package endpoint

import "github.com/sqdron/go-squad/endpoint/http"

type transport struct {
	provider interface{}
}

type ITransport interface {
	Options() []interface{}
	//Decode(data interface{}) Message
	//Encode(m Message) interface{}
}

func NewTransport(provider interface{}) *transport {
	return &transport{provider:provider}
}

func (t *transport) Options() []interface{} {
	return t.provider.(ITransport).Options()
}

func Http() *transport{
	return NewTransport(&http.HttpTransport{})
}