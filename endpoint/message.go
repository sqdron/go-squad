package endpoint

type Message struct {
	ID       interface{}
	Responce string
	Payload  interface{}
}

//func CreateMessage(data interface{}) *Message {
//	return &Message{ID:util.GenerateString(10), Payload: json.Marshal(data), Responce: "resp_" + util.GenerateString(5)}
//}
//
//func CreateResponceMessage(source *Message, data interface{})*Message {
//	return &Message{ID:util.GenerateString(10), Payload: json.Marshal(data)}
//}
