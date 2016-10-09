package connect

import (
	"encoding/json"
	"github.com/nats-io/nats"
	"github.com/sqdron/squad/middleware"
	"reflect"
)

func CreateEncoderMiddleware(dataObj interface{}) middleware.Middleware {
	dataType := reflect.TypeOf(dataObj)
	return func(next middleware.IAction) middleware.IAction {
		return middleware.ActionFunc(func(r interface{}) interface{} {
			msg := r.(*nats.Msg)
			var oPtr reflect.Value
			if dataType.Kind() != reflect.Ptr {
				oPtr = reflect.New(dataType)
			} else {
				oPtr = reflect.New(dataType.Elem())
			}
			e := json.Unmarshal(msg.Data, oPtr.Interface())
			if e != nil {
				panic(e)
			}
			return next.Apply(oPtr.Elem().Interface())
		})
	}
}
