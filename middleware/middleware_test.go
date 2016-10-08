package middleware

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

type testContext struct {
}

func (a *testContext) Apply(i interface{}) interface{} {
	fmt.Println(i)
	return i
}

func Test_Middleware(t *testing.T) {
	Convey("Create middleware", t, func() {
		var logger Middleware = func(next IAction) IAction {
			return ActionFunc(func(r interface{}) interface{} {
				return next.Apply(r)
			})
		}

		var double Middleware = func(next IAction) IAction {
			return ActionFunc(func(r interface{}) interface{} {
				res := r.(int) * 2
				return next.Apply(res)
			})
		}
		var addThree Middleware = func(next IAction) IAction {
			return ActionFunc(func(r interface{}) interface{} {
				res := r.(int) + 3
				return next.Apply(res)
			})
		}
		act := &testContext{}
		m := ApplyMiddleware(logger, double, addThree)
		r := m(act).Apply(5)
		So(r, ShouldNotBeNil)
		So(r.(int), ShouldEqual, 13)

	})
}
