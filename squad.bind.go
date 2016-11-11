package squad

import (
	"github.com/sqdron/squad/connect"
	"time"
)

type squadBind struct {
	transport connect.ITransport
	//actions   map[string]interface{}
	//policyMap map[string]*policy.Policy
	//exit      chan bool
}

type InvokerFunc func(request interface{}) (interface{}, error)

type SquadBinder struct {
	invoker func(request interface{}) (interface{}, error)
	//bind func()
}

type ISquadBinder interface {
	setInvoker(i InvokerFunc)
	Invoke(request interface{}) (interface{}, error)
	//binder(func())
}
//
type ISquadBind interface {
	ToController(path string, b ISquadBinder) interface{}
}

func CreateBind(transport connect.ITransport) *squadBind {
	return &squadBind{transport:transport}
}

func (s *squadBind) ToController(path string, b ISquadBinder) interface{} {
	b.setInvoker(func(request interface{}) (interface{}, error){
		return s.transport.RequestSync(path, request, 3 * time.Second)
	})
	return b
}

func (s *SquadBinder) setInvoker(i InvokerFunc) {
	s.invoker = i
}

func (s *SquadBinder) Invoke(request interface{}) (interface{}, error) {
	return s.invoker(request)
}
