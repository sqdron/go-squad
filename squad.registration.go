package squad

import (
	"fmt"
	"github.com/sqdron/squad/hub"
)

//register - just send service specification to hub(vesrion, name, api enpoints...). No need in error handling on any business logic.
func (s *squad) register() error {
	fmt.Println("Register service")
	s.Connect.Request("squad.register", nil, func() {})
	h := hub.HubClient(s.Connect)
	return h.Register(s.Spec)
}

