package squad

import (
	"github.com/sqdron/squad/configurator"
	"github.com/sqdron/squad/connect"
	"fmt"
)

type ISquad interface {
	Activate()
}

type ActivationOptions struct {
	Name string
	Hub  string
}

// squad is an isolated micro-service unit. Nobody knows about his existence but it knows about the entire world.
type squad struct {
	options *ActivationOptions
	modules map[string]interface{}
	Connect connect.ITransport
	Api     ISquadAPI
}

func (s *squad) use(module interface{}) {
}

func (s *squad) Options() *ActivationOptions {
	return s.options
}

func Client() *squad {

	opts := &ActivationOptions{}
	cfg := configurator.New()
	cfg.ReadFlags(opts)
	fmt.Println(opts.Hub)

	client := &squad{options: opts}
	client.Api = CreateApi()
	return client
}
