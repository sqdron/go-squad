package service

type Action struct {
	Name        string
	Description string
	InputType   string
	OutputType  string
	Input       []ParamInfo
	Output      []ParamInfo
}

type ParamInfo struct {
	Name string
	Type string
}

//ServiceSpecification - is a description of micro-service application. Goes to Hub
type Specification struct {
	Name    string
	Version string
	Actions []Action
}
