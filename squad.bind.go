package squad

import (
	"encoding/json"
	"fmt"
	"github.com/sqdron/squad/connect"
	"time"
)

type squadBind struct {
	transport connect.ITransport
	//actions   map[string]interface{}
	//policyMap map[string]*policy.Policy
	//exit      chan bool
}

type InvokerFunc func(action string, request interface{}) (interface{}, error)

type SquadBinder struct {
	invoker func(action string, request interface{}) (interface{}, error)
	//bind func()
}

type ISquadBinder interface {
	setInvoker(i InvokerFunc)
	Invoke(action string, request interface{}, result interface{}) error
	//binder(func())
}

//
type ISquadBind interface {
	ToController(path string, b ISquadBinder) interface{}
}

func CreateBind(transport connect.ITransport) *squadBind {
	return &squadBind{transport: transport}
}

func (s *squadBind) ToController(path string, b ISquadBinder) interface{} {
	b.setInvoker(func(action string, request interface{}) (interface{}, error) {
		subject := path + "." + action
		res, err := s.transport.RequestSync(subject, request, 3*time.Second)
		fmt.Println("requset error", err)
		return res, err
	})
	return b
}

func (s *SquadBinder) setInvoker(i InvokerFunc) {
	s.invoker = i
}

func (s *SquadBinder) Invoke(action string, request interface{}, result interface{}) error {
	res, err := s.invoker(action, request)
	json.Unmarshal(res.([]byte), &result)
	return err
}
