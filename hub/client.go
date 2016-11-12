package hub

import (
	"github.com/pkg/errors"
	"github.com/sqdron/squad/connect"
	"github.com/sqdron/squad/service"
	"sync"
	"time"
)

type hubClient struct {
	connect connect.ITransport
}

func HubClient(c connect.ITransport) *hubClient {
	return &hubClient{connect: c}
}

func (h *hubClient) Register(spec *service.Specification) error {
	_, err := h.connect.RequestSync(HUB_REGISTER_ENDPOINT, spec, 3*time.Second)
	return err
}

func (h *hubClient) Activate(name string) (*service.Instruction, error) {
	wg := sync.WaitGroup{}
	wg.Add(1)
	var res *service.Instruction = nil
	var resError error = nil
	go func() {
		resError = h.connect.Request(HUB_ACTIVATE_ENDPOINT, &service.Specification{Name: name}, func(i *service.Instruction) {
			res = i
			wg.Done()
		})
		<-time.Tick(3 * time.Second)
		resError = errors.New("Timeout")
	}()
	wg.Wait()
	return res, resError
}
