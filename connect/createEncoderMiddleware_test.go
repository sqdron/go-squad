package connect

import (
	"fmt"
	"github.com/nats-io/nats"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/sqdron/squad/middleware"
	"testing"
	"encoding/json"
)

type testRequest struct {
	Name  string
	Value int
}

func Test_Middleware(t *testing.T) {
	Convey("Create middleware", t, func() {
		//enc := &JsonEncoder{}
		data := &testRequest{Name: "testrequest", Value: 5}
		bytes, _ := json.Marshal(data)
		msg := &nats.Msg{Subject: "test", Data: bytes}
		action := func(r testRequest) int {
			fmt.Println("Action handler")
			return 7
		}

		mware := CreateEncoderMiddleware(&testRequest{})
		m := middleware.ApplyMiddleware(mware)
		ctx := &requestContext{action: action}
		r := m(ctx).Apply(msg)
		So(r, ShouldNotBeNil)
		So(r.(int), ShouldEqual, 7)
	})
}
