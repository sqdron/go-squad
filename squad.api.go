package squad

import (
	"fmt"
)

type Action func(request interface{}) interface{}
type ApiFunc func(action interface{})

type AcceptAction func(ApiFunc)

type api struct {
	actions map[string]interface{}
}

func (s *squad) initAPI() {
	s.api = &api{actions:make(map[string]interface{})}
}

func (s *squad) Api(path string) ApiFunc {
	fmt.Println("Registering route: " + path)
	_, exist := s.api.actions[path];
	if (exist) {
		panic("Api route already exist: " + path)
	}

	return func(action interface{}) {
		fmt.Println("Registering action: ")
		s.api.actions[path] = action
		listen := s.Endpoint.Listen(path)
		fmt.Println("starting goroutin: " + path)
		go func() {
			for {
				message := <-listen
				fmt.Println(message)
			}
		}()
	}
}

func (af ApiFunc) Action(action interface{}) {
	af(action)
}


