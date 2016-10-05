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
		s.api.actions[path] = action
		fmt.Println("Listen...")
		s.Endpoint.Listen(path, action)
	}
}

func (af ApiFunc) Action(action interface{}) {
	af(action)
}
