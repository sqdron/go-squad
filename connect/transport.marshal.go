package connect

import (
	"encoding/json"
	"golang.org/x/crypto/openpgp/errors"
	"log"
	"reflect"
)

func marshalMessage(obj interface{}) ([]byte, error) {
	if obj == nil {
		return nil, nil
	}

	switch obj.(type) {
	default:
		return json.Marshal(obj)
	case []byte:
		return obj.([]byte), nil
	}
}

func unmarshalMessage(subject string, data []byte, action interface{}) (interface{}, error) {
	actionType := reflect.TypeOf(action)
	arguments := []reflect.Value{}
	numOut := actionType.NumOut()

	if actionType.NumIn() > 1 {
		return nil, errors.SignatureError("Too many input arguments for action " + subject)
	}
	if numOut > 2 {
		return nil, errors.SignatureError("Too many outputs for action " + subject)
	}
	log.Println("")
	if actionType.NumIn() == 1 {
		inType := actionType.In(0)
		var oPtr reflect.Value
		switch inType.Kind() {
		case reflect.Ptr:
			oPtr = reflect.New(inType.Elem())
			arguments = append(arguments, oPtr)
		default:
			oPtr = reflect.New(inType)
			arguments = append(arguments, oPtr.Elem())
		}

		e := json.Unmarshal(data, oPtr.Interface())
		log.Println(oPtr.Elem().Interface())
		if e != nil {
			return nil, e
		}
	}

	result := reflect.ValueOf(action).Call(arguments)
	switch numOut {
	case 1:
		return result[0].Interface(), nil
	case 2:
		var e error = nil
		if result[1].Interface() != nil {
			e = result[1].Interface().(error)
		}
		return result[0].Interface(), e
	default:
		return nil, nil
	}
}
