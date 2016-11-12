package squad

import (
	"fmt"
	"github.com/sqdron/squad/hub"
)

func (s *squad) activate() error {
	fmt.Println("Activate...")
	h := hub.HubClient(s.Connect)
	instruction, err := h.Activate(s.Name)
	if err != nil {
		fmt.Errorf("Activation error", err)
		return err
	}
	fmt.Println("Instructions...")
	s.Instruction = instruction
	fmt.Println(s.Instruction)
	return nil
}
