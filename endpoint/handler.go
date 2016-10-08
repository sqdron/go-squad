package endpoint

import (
	"errors"
	"fmt"
	"reflect"
)

type ApiFunc func(*Message)

type ApiHandler interface {
	ServeMessage(*Message)
}

func (m *Message) Apply(action interface{}, data interface{}) interface{} {
	content := acceptInputArguments(data, action)
	var res = reflect.ValueOf(action).Call(content)[0]
	return res.Interface()
}

func acceptInputArguments(payload interface{}, action interface{}) []reflect.Value {
	payloadValue := reflect.ValueOf(payload)
	actionValue := reflect.ValueOf(action)

	result := []reflect.Value{}

	actionType := actionValue.Type()
	if actionType.NumIn() == 0 {
		return result
	}
	if !payloadValue.IsValid() {
		panic("Payload is invalid!")
	}
	if actionType.NumIn() > 1 {
		panic("Parameters count mismatch!")
	}

	params := payload.(map[string]interface{})
	if actionType.In(0).Kind() == reflect.Struct {
		instance := reflect.New(actionType.In(0))
		FillStruct(instance, params)
		return append(result, instance.Elem())
	}
	return result
}

func FillStruct(s reflect.Value, m map[string]interface{}) error {
	for k, v := range m {
		err := SetField(s, k, v)
		if err != nil {
			return err
		}
	}
	return nil
}

func SetField(obj reflect.Value, name string, value interface{}) error {
	structValue := obj.Elem()
	structFieldValue := structValue.FieldByName(name)

	if !structFieldValue.IsValid() {
		return fmt.Errorf("No such field: %s in obj", name)
	}

	if !structFieldValue.CanSet() {
		return fmt.Errorf("Cannot set %s field value", name)
	}

	structFieldType := structFieldValue.Type()
	val := reflect.ValueOf(value)
	if structFieldType != val.Type() {
		return errors.New("Provided value type didn't match obj field type")
	}

	structFieldValue.Set(val)
	return nil
}
