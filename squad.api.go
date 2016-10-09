package squad

import (
	"github.com/sqdron/squad/activation"
	"github.com/sqdron/squad/connect"
	"reflect"
)

type squadApi struct {
	transport connect.ITransport
	actions   map[string]interface{}
}

type EndpointFunc func(interface{})
type ActionFunc func(action interface{})

type ISquadAPI interface {
	Route(path string) ActionFunc
	Request(path string, message interface{}, cb interface{}) error
	//RequestSync(path string, message interface{}) interface{}

	getMetadata() []activation.ActionMeta
	start(i *activation.ServiceInfo)
}

func (api *squadApi) Route(path string) ActionFunc {
	return func(action interface{}) {
		api.actions[path] = action
	}
}

func (af ActionFunc) Action(action interface{}) {
	af(action)
}

func (api *squadApi) start(i *activation.ServiceInfo) {
	api.transport = connect.NewTransport(i.Endpoint)
	for path, action := range api.actions {
		func(info *activation.ServiceInfo) {
			api.transport.QueueSubscribe(path, info.Group, action)
		}(i)
	}
}

func (api *squadApi) Request(path string, message interface{}, cb interface{}) error {
	return api.transport.Request(path, message, cb)
}

//func (api *squadApi) RequestSync(path string, message interface{}) interface{} {
//	return api.transport.RequestSync(path, message)
//}

func CreateApi() *squadApi {
	return &squadApi{actions: make(map[string]interface{})}
}

func (api *squadApi) getMetadata() []activation.ActionMeta {
	result := []activation.ActionMeta{}
	for route, action := range api.actions {
		am := activation.ActionMeta{Name: route, Input: []activation.ParamMeta{}, Output: []activation.ParamMeta{}}

		aType := reflect.TypeOf(action)
		for i := 0; i < aType.NumIn(); i++ {
			inType := aType.In(i)
			am.InputType = inType.Name()
			for f := 0; f < inType.NumField(); f++ {
				field := inType.Field(f)
				am.Input = append(am.Input, activation.ParamMeta{Name: field.Name, Type: field.Type.Name()})
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
