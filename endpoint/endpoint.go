package endpoint

import (
	"context"
)

type MessageHandler func(ctx context.Context, request Message) (response Message, err error)

type EndPoint struct {
	Transport interface{}
	publish MessageHandler
	listen MessageHandler

}

type IEndpoint interface{
	Publish() chan <- Message
	Listen() <- chan Message
}



//
//func (s *EndPoint) Use(m Middleware) EndPoint{
//	return s.transport
//}


//
//type Endpoint func(ctx context.Context, request interface{}) (response interface{}, err error)
//
type Middleware func(MessageHandler) MessageHandler


//
//func (m Endpoint) Compose(endpoint Endpoint) Endpoint {
//	return func(ctx interface{}) interface{} {
//		return m(endpoint(ctx))
//	}
//}
//
////ComposeAll middleware
//func (m Endpoint) ComposeAll(endpoints ...Endpoint) Endpoint {
//	return func(ctx interface{}) interface{} {
//		res := m
//		for i := 0; i < len(endpoints); i++ {
//			res = res.Compose(endpoints[i])
//		}
//		return res(ctx)
//	}
//}
//
////New instance of the transport protocol
//func New(provider interface{}) transport{
//	return &transport{provider:provider}
//}
//
////Run real implementation
//func(t *transport) Run(){
//	t.provider.(ITransport).Run();
//}