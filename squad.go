package squad

import (
	"github.com/sqdron/squad/configurator"
	"github.com/sqdron/squad/connect"
)

type ISquad interface {
	Activate()
}

type ActivationOptions struct {
	ApplicationID  string `json:"app_id"`
	ApplicationHub string `json:"app_hub"`
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

func Client(config ...string) *squad {
	connectionConfig := "connect.json"
	if len(config) > 0 {
		connectionConfig = config[0]
	}
	opts := &ActivationOptions{}
	cfg := configurator.New()
	cfg.ReadFromFile(connectionConfig, opts)
	client := &squad{options: opts}
	client.Api = CreateApi()
	return client
}
