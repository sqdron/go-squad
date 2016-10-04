package endpoint

type IEndpoint interface {
	Publish(subject string) chan <- *Message
	Listen(subject string) <- chan *Message
	Close()
}

type Endpoint struct {
	provider interface{}
}

func NewEndpoint(provider interface{}) *Endpoint {
	return &Endpoint{provider:provider}
}

func (t *Endpoint) Publish(subject string) chan <- *Message {
	return t.provider.(IEndpoint).Publish(subject)
}

func (t *Endpoint) Listen(subject string) <- chan *Message {
	return t.provider.(IEndpoint).Listen(subject)
}
func (t *Endpoint) Close() {
	t.provider.(IEndpoint).Close()
}
