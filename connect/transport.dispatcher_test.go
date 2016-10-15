package connect

import (
	"errors"
	"fmt"
	"github.com/smartystreets/assertions/should"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

type testMessage struct {
	Param1 string
	Param2 int
}

func Test_Dispatch_Message_General(t *testing.T) {
	Convey("Dispatch Message: ", t, func() {
		Convey("General Case", func() {
			res, e := applyMessage("subj1", prepareMessageData(), func(r testMessage) (interface{}, error) {
				return r.Param2, nil
			})
			So(e, should.BeNil)
			So(res, should.Equal, 42)
		})

		Convey("Empty result", func() {
			res, e := applyMessage("subj1", prepareMessageData(), func(r testMessage) {
				fmt.Println(r)
			})
			So(e, should.BeNil)
			So(res, should.BeNil)
		})

		Convey("Empty Message", func() {
			res, e := applyMessage("subj1", prepareMessageData(), func() int {
				return 42
			})
			So(e, should.BeNil)
			So(res, should.Equal, 42)
		})

		Convey("General Case With Error", func() {
			_, e := applyMessage("subj1", prepareMessageData(), func(r testMessage) (interface{}, error) {
				return nil, errors.New("Some error")
			})
			fmt.Println(e)
			So(e.Error(), should.Equal, "Some error")
		})

		Convey("Too many inputs", func() {
			_, e := applyMessage("subj1", prepareMessageData(), func(r testMessage, extra string) (interface{}, error) {
				return 5, errors.New("Some error")
			})
			So(e, should.NotBeNil)
		})

		Convey("Too many outputs", func() {
			_, e := applyMessage("subj1", prepareMessageData(), func(r testMessage, extra string) (interface{}, error, int) {
				return 5, errors.New("Some error"), 1
			})
			So(e, should.NotBeNil)
		})

		Convey("Message is string", func() {
			m, _ := marshalMessage("test message")
			res, e := applyMessage("subj1", m, func(r string) (string, error) {
				return r, errors.New("Some error")
			})
			So(e, should.NotBeNil)
			So(res, should.Equal, "test message")
		})
	})
}

func prepareMessageData() []byte {
	res, _ := marshalMessage(&testMessage{
		Param1: "test data",
		Param2: 42})
	return res
}
