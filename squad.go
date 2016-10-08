package squad

import (
	"github.com/sqdron/squad/configurator"
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
	applicationId string
	endpoint      string
	modules       map[string]interface{}
	Api           ISquadAPI
}

func (s *squad) use(module interface{}) {

}

func Client(config ...string) *squad {
	connectionConfig := "config.json"
	if (len(config) > 0) {
		connectionConfig = config[0]
	}
	opts := &ActivationOptions{}
	cfg := configurator.New()
	cfg.ReadFromFile(connectionConfig, opts)
	client := &squad{applicationId: opts.ApplicationID, endpoint: opts.ApplicationHub}
	client.Api = CreateApi(opts.ApplicationHub)
	return client
}
