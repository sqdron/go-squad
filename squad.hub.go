package squad

import (
	"github.com/sqdron/squad/configurator"
	"github.com/sqdron/squad/connect"
	"github.com/sqdron/squad/hub"
	"fmt"
	"os"
	"os/signal"
)

type hubServer struct {
	Api     hub.IHubAPI
	Options *ClientOptions
	Run     func()
}

func Hub(options ...interface{}) *hubServer {
	opts := &ClientOptions{}
	cfg := configurator.New()
	cfg.MapOptions(opts)
	for _, o := range options {
		cfg.MapOptions(o)
	}
	cfg.ReadOptions()
	t := connect.NatsTransport(opts.Hub)
	hubApi := hub.HubApi(t)
	client := &hubServer{Api:hubApi}
	client.Options = opts
	client.Run = func() {
		fmt.Println("Run...")
		if (opts.Specification != nil && *opts.Specification == true) {
		} else {
			hubApi.Start()
			h := hub.HubClient(t)
			_, err := h.Activate("squad.hub")
			if (err != nil) {
				fmt.Errorf("Activation error", err)
			}
			<- client.start()
		}
	}
	return client
}

//TODO: refactor code duplication
func (s *hubServer) start() <- chan bool {
	exit := make(chan bool, 1)
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	//wg := sync.WaitGroup{}
	//defer wg.Wait()
	//wg.Add(1)
	go func() {
		for {
			select {
			case <-c:
				fmt.Println("Interrupting...")
			//wg.Done()
			//	s.Api.stop()
				exit <- true
				os.Exit(1)
			}
		}
	}()
	//wg.Wait()
	return exit
}
