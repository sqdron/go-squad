package squad

import (
	"github.com/sqdron/squad/activation"
	"github.com/sqdron/squad/connect"
	"reflect"
)

type squadApi struct {
	transport connect.ITransport
	actions   map[string]interface{}
	requests  map[string]interface{}
}

type EndpointFunc func(interface{})
type ActionFunc func(action interface{})
type RequestFunc func(message interface{})

type ISquadAPI interface {
	Route(path string) ActionFunc
	Request(path string) RequestFunc
	//RequestSync(path string, message interface{}) interface{}

	getMetadata() []activation.ActionMeta
	start(i *activation.ServiceInfo)
}

func (api *squadApi) Route(path string) ActionFunc {
	return func(action interface{}) {
		api.actions[path] = action
	}
}

func (api *squadApi) Request(path string) RequestFunc {
	return func(request interface{}) {
		api.requests[path] = request
	}
}

func (af ActionFunc) Action(action interface{}) {
	af(action)
}

func (rf RequestFunc) Submit(message interface{}) {
	rf(message)
}

func (api *squadApi) start(i *activation.ServiceInfo) {
	api.transport = connect.NatsTransport(i.Endpoint)
	for path, action := range api.actions {
		func(info *activation.ServiceInfo) {
			api.transport.QueueSubscribe(path, info.Group, action)
		}(i)
	}
	for path, request := range api.requests {
		func(info *activation.ServiceInfo) {
			api.transport.Publish(path, request)
		}(i)
	}
}

func CreateApi() *squadApi {
	return &squadApi{actions: make(map[string]interface{}), requests: make(map[string]interface{})}
}

func (api *squadApi) getMetadata() []activation.ActionMeta {
	result := []activation.ActionMeta{}
	for route, action := range api.actions {
		am := activation.ActionMeta{Name: route, Input: []activation.ParamMeta{}, Output: []activation.ParamMeta{}}

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
					am.Input = append(am.Input, activation.ParamMeta{Name: field.Name, Type: field.Type.Name()})
				}
			}
		}

		for i := 0; i < aType.NumOut(); i++ {
			outType := aType.Out(i)
			am.OutputType = outType.Name()
			if outType.Kind() == reflect.Struct {
				for f := 0; f < outType.NumField(); f++ {
					field := outType.Field(f)
					am.Output = append(am.Output, activation.ParamMeta{Name: field.Name, Type: field.Type.Name()})
				}
			}
		}
		result = append(result, am)
	}
	return result
}
