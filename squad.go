package squad

import (
	"github.com/sqdron/squad/endpoint"
	"github.com/sqdron/squad/endpoint/nats"
)

type ISquad interface {
	Activate()
}

//type Options struct {
//	Id      string
//	Name    string
//	Version string
//	Secret	string
//}

// squad is an isolated micro-service unit. Nobody knows about his existence but it knows about the entire world.
type squad struct {
	ApplicationId string
	api           *api
	Endpoint      endpoint.IEndpoint
}

func Client(url string, appId string) *squad {
	client := &squad{ApplicationId:appId, Endpoint:nats.NatsEndpoint(url)}
	client.initAPI()
	return client
}
