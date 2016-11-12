package squad

import (
	"github.com/sqdron/squad/connect"
	"reflect"
	"github.com/sqdron/squad/policy"
	"github.com/sqdron/squad/service"
	"fmt"
	"strings"
	"time"
)

type squadApi struct {
	transport connect.ITransport
	actions   map[string]interface{}
	policyMap map[string]*policy.Policy
	exit      chan bool
}

type EndpointFunc func(interface{}) interface{}
type ActionFunc func(string, interface{}) IActionPolicy
type RequestFunc func(message interface{})

type ISquadAPI interface {
	Controller(controller interface{})

	Remote(controller interface{}, aptr interface{})
	//RemoteController(val interface{}) (interface{}, error)

	//RemoteAction(controller interface{}, aptr interface{}) error

	Request(resource string, message interface{}) (interface{}, error)
	getActions() []service.Action
	start(i *service.Instruction)
	stop()
}

func (api *squadApi) GetAction(resource string) EndpointFunc {
	return func(message interface{}) interface{} {
		return nil
		//return api.transport.RequestSync(resource, message, 3 * time.Second)
	}
}

//func (api *squadApi) Route(controller string) ActionFunc {
//	return func(name string, action interface{}) IActionPolicy {
//		route := fmt.Sprintf("%s.%s", controller, name)
//		api.actions[route] = action
//		return ActionPolicy(route)
//	}
//}


func (api *squadApi) Controller(controller interface{}) {
	cType := reflect.TypeOf(controller)
	vType := reflect.ValueOf(controller)
	for i := 0; i < cType.NumMethod(); i++ {
		m := cType.Method(i)
		route := strings.ToLower(fmt.Sprintf("%s.%s", cType.Elem().Name(), m.Name))
		api.actions[route] = vType.Method(i).Interface()
	}
}

func (api *squadApi) Remote(controller interface{}, aptr interface{}) {
	wrap := func(in []reflect.Value) []reflect.Value {

		fmt.Println("Bingo!!!")
		var errorType = reflect.TypeOf(make([]error, 1)).Elem()
		return []reflect.Value{reflect.Zero(reflect.TypeOf("")), reflect.Zero(errorType)}
	}
	fn := reflect.ValueOf(aptr).Elem()
	fmt.Println("fn.Type()", fn.Type())
	// Make a function of the right type.
	v := reflect.MakeFunc(fn.Type(), wrap)

	// Assign it to the value fn represents.
	fn.Set(v)
}

func (af ActionFunc) Action(name string, action interface{}) IActionPolicy {
	return af(name, action)
}

func (rf RequestFunc) Submit(message interface{}) {
	rf(message)
}

func (api *squadApi) Request(resource string, message interface{}) (interface{}, error) {
	return api.transport.RequestSync(resource, message, 3 * time.Second)
}

func (api *squadApi) start(i *service.Instruction) {
	for path, action := range api.actions {
		api.transport.QueueSubscribe(path, i.Group, action)
	}
}

func (api *squadApi) stop() {
	//api.exit <- true
}

func CreateApi(transport connect.ITransport) *squadApi {
	return &squadApi{actions: make(map[string]interface{}),
		policyMap: make(map[string]*policy.Policy),
		exit:make(chan bool), transport:transport}
}

func (api *squadApi) getActions() []service.Action {
	result := []service.Action{}
	for route, action := range api.actions {
		am := service.Action{Name: route, Input: []service.ParamInfo{}, Output: []service.ParamInfo{}}

		aType := reflect.TypeOf(action)
		for i := 0; i < aType.NumIn(); i++ {
			inType := aType.In(i)
			if inType.Kind() == reflect.Ptr {
				inType = inType.Elem()
			}
			am.InputType = inType.Name()
			if inType.Kind() == reflect.Struct {
				for f := 0; f < inType.NumField(); f++ {
					field := inType.Field(f)
					am.Input = append(am.Input, service.ParamInfo{Name: field.Name, Type: field.Type.Name()})
				}
			}
		}

		for i := 0; i < aType.NumOut(); i++ {
			outType := aType.Out(i)
			am.OutputType = outType.Name()
			if outType.Kind() == reflect.Struct {
				for f := 0; f < outType.NumField(); f++ {
					field := outType.Field(f)
					am.Output = append(am.Output, service.ParamInfo{Name: field.Name, Type: field.Type.Name()})
				}
			}
		}
		result = append(result, am)
	}
	return result
}
