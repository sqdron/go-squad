package squad

import (
	"github.com/sqdron/squad/endpoint"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"github.com/sqdron/squad/endpoint/nats"
	"crypto/rand"
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
	endpoint      endpoint.IEndpoint
}

func Client(url string, appId string) *squad {
	return &squad{ApplicationId:appId, endpoint:nats.NatsEndpoint(url)}
}

func randonString(strSize int) string {
	var dictionary string = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

	var bytes = make([]byte, strSize)
	rand.Read(bytes)
	for k, v := range bytes {
		bytes[k] = dictionary[v % byte(len(dictionary))]
	}
	return string(bytes)
}

