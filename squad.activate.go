package squad

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"github.com/sqdron/squad/activation"
)

func (s *squad) Activate() {
	fmt.Println("Activation...")

	act := activation.RequestActivation{ID:s.applicationId, Actions:s.Api.getMetadata()}
	restartApi := func(info activation.ServiceInfo) bool {
		fmt.Println("Restart requested")
		fmt.Println(info)
		s.Api.start(&info)
		return true
	}
	s.Api.Request("activate", act, restartApi)
	s.start()
}

func (s *squad) RunDetached() {
	fmt.Println("Running detached...")
	s.Api.start(&activation.ServiceInfo{Group:""})
	s.start()
}

func (s *squad) start() {
	fmt.Println("Starting...")
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

