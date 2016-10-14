package squad

import (
	"fmt"
	"github.com/sqdron/squad/activation"
	"github.com/sqdron/squad/connect"
	"log"
	"os"
	"os/signal"
	"sync"
)

func (s *squad) Activate(cb ...func(activation.ServiceInfo)) {
	s.Connect = connect.NatsTransport(s.options.Hub)

	act := activation.RequestActivation{ID: s.options.Name, Actions: s.Api.getMetadata()}
	restartApi := func(info activation.ServiceInfo) bool {
		log.Println("Activation complited!")
		s.Api.start(&info)
		if len(cb) > 0 {
			cb[0](info)
		}
		return true
	}
	s.Connect.Request("activate", act, restartApi)
	s.start()
}

func (s *squad) RunDetached() {
	fmt.Println("Running detached...")
	s.Api.start(&activation.ServiceInfo{Group: "", Endpoint: s.options.Hub})
	s.start()
}

func (s *squad) start() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	wg := sync.WaitGroup{}
	defer wg.Wait()
	wg.Add(1)
	go func() {
		for {
			select {
			case <-c:
				fmt.Println("Interrupting...")
				wg.Done()
				os.Exit(1)
			}
		}
	}()
	wg.Wait()
}
