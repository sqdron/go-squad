package squad

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
)

func (s *squad) Activate() {
	fmt.Println("Activation...")

	//message := &endpoint.Message{}
	//message.Content = make(map[string]interface{})
	//message.Content["responce"] = randonString(10)
	//message.Content["app-id"] = s.ApplicationId;
	//
	//s.endpoint.Publish("activate") <- message

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

