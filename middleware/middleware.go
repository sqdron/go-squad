package middleware

import "fmt"

type IAction interface {
	Apply(request interface{}) interface{}
}

type DefaultContext struct {
}

func (d *DefaultContext) Apply(request interface{}) interface{} {
	fmt.Println("Default middleware")
	return request
}

type ActionFunc func(request interface{}) interface{}

func (f ActionFunc) Apply(r interface{}) interface{} {
	return f(r)
}

type Middleware func(next IAction) IAction

func ApplyMiddleware(middleware ...Middleware) Middleware {
	return func(action IAction) IAction {
		if action == nil {
			action = &DefaultContext{}
		}

		size := len(middleware) - 1
		for i := range middleware {
			action = middleware[size-i](action)
		}
		return action
	}
}
