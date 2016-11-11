package service

import "github.com/sqdron/squad/policy"

//ServiceInstruction - is an instruction for application. Get it before run
type Instruction struct {
	ID       string
	Group    string
	Endpoint string
	Policy   []*policy.Policy
}
