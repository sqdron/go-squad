package endpoint

type ApiFunc func(*Message)

type ApiHandler interface {
	ServeMessage(/*ResponseWriter,*/ *Message)
}

//func Middleware(h ApiHandler) ApiHandler {
//	return ApiHandlerFunc(func(m *Message) {
//		h.ServeMessage(m)
//	})
//}