package connect

type Message struct {
	ID      string
	Subject string
	Reply   string
	Data    interface{}
}
