package endpoint

type IEndpoint interface {
	Listen(subject string, handler interface{})
	//Publish(subject string, data interface{})
	Request(subject string, data interface{}) <-chan interface{}
	Close()

	//Listen(subject string) <- chan interface{}
	//Publish(subject string, data interface{})
	//Request(subject string, data interface{}) <- chan interface{}
	//Close()
}

type Endpoint struct {
	provider interface{}
}

func NewEndpoint(provider interface{}) *Endpoint {
	return &Endpoint{provider: provider}
}

func (t *Endpoint) Listen(subject string, handler interface{}) {
	t.provider.(IEndpoint).Listen(subject, handler)
}

//func (t *Endpoint) Publish(subject string, data interface{}){
//	t.provider.(IEndpoint).Publish(subject, data)
//}

func (t *Endpoint) Request(subject string, data interface{}) <-chan interface{} {
	return t.provider.(IEndpoint).Request(subject, data)
}

func (t *Endpoint) Close() {
	t.provider.(IEndpoint).Close()
}
