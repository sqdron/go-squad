package endpoint

type IEndpoint interface {
	Publish(subject string) chan <- *Message
	Listen(subject string) <- chan *Message
	Close()
}

type Transport struct {
	provider interface{}
}

func NewEndpoint(provider interface{}) *Transport {
	return &Transport{provider:provider}
}

func (t *Transport) Publish(subject string) chan <- *Message {
	return t.provider.(IEndpoint).Publish(subject)
}

func (t *Transport) Listen(subject string) <- chan *Message {
	return t.provider.(IEndpoint).Listen(subject)
}
func (t *Transport) Close() {
	t.provider.(IEndpoint).Close()
}

//////
//////type Endpoint func(ctx context.Context, request interface{}) (response interface{}, err error)
//////
////type Middleware func(MessageHandler) MessageHandler
//
////
////func (m Endpoint) Compose(endpoint Endpoint) Endpoint {
////	return func(ctx interface{}) interface{} {
////		return m(endpoint(ctx))
////	}
////}
////
//////ComposeAll middleware
////func (m Endpoint) ComposeAll(endpoints ...Endpoint) Endpoint {
////	return func(ctx interface{}) interface{} {
////		res := m
////		for i := 0; i < len(endpoints); i++ {
////			res = res.Compose(endpoints[i])
////		}
////		return res(ctx)
////	}
////}
////
//////New instance of the transport protocol
////func New(provider interface{}) transport{
////	return &transport{provider:provider}
////}
////
//////Run real implementation
////func(t *transport) Run(){
////	t.provider.(ITransport).Run();
////}
