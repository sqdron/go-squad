package hub

import "github.com/sqdron/squad/connect"

type hubApi struct {
	transport connect.ITransport
	register  interface{}
	activate  interface{}
}

type IHubAPI interface {
	Register(action interface{})
	Activate(action interface{})
}

func (h *hubApi) Register(action interface{}) {
	h.register = action
}

func (h *hubApi) Activate(action interface{}) {
	h.activate = action
}

func (api *hubApi) Start() {
	api.transport.QueueSubscribe(HUB_REGISTER_ENDPOINT, "hub", api.register)
	api.transport.QueueSubscribe(HUB_ACTIVATE_ENDPOINT, "hub", api.activate)
}

func HubApi(t connect.ITransport) *hubApi {
	return &hubApi{transport: t}
}
