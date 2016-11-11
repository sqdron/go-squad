package squad

import (
	"github.com/sqdron/squad/configurator"
	"github.com/sqdron/squad/connect"
	"github.com/sqdron/squad/service"
	"fmt"
	"os"
	"os/signal"
)

type ISquad interface {
	Run()
}

type ClientOptions struct {
	ID            string
	Hub           string
	Specification *bool
}

// squad is an isolated micro-service unit. Nobody knows about his existence but it knows about the entire world.
type squad struct {
	ID          string
	Name        string
	Options     *ClientOptions
	Connect     connect.ITransport
	Api         ISquadAPI
	Bind        ISquadBind
	Spec        *service.Specification
	Instruction *service.Instruction
}

func Client(name string, options ...interface{}) *squad {
	opts := &ClientOptions{}
	cfg := configurator.New()
	cfg.MapOptions(opts)
	for _, o := range options {
		cfg.MapOptions(o)
	}
	cfg.ReadOptions()
	client := &squad{Options: opts}
	client.Name = name
	client.ID = opts.ID
	client.Connect = connect.NatsTransport(opts.Hub)
	client.Api = CreateApi(client.Connect)
	client.Bind = CreateBind(client.Connect)
	return client
}

func (s *squad) Run() <- chan bool {
	fmt.Println("Run...")
	if (s.Options.Specification != nil && *s.Options.Specification == true) {
		s.register()
	} else {
		err := s.activate()
		if (err == nil) {
			s.Api.start(s.Instruction)
			return s.start()
		}
	}
	return nil
}

func (s *squad) start() <- chan bool {
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
				s.Api.stop()
				exit <- true
				os.Exit(1)
			}
		}
	}()
	//wg.Wait()
	return exit
}
