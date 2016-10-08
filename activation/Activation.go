package activation

type RequestActivation struct {
	ID      string
	Actions []ActionMeta
}

type ResponceActivation struct {
	ID string
	Group string
}

type ActionMeta struct {
	Name        string
	Description string
	InputType   string
	OutputType  string
	Input       []ParamMeta
	Output      []ParamMeta
}

type ParamMeta struct {
	Name string
	Type string
}

type ServiceInfo struct {
	Endpoint string
	Group    string
	Options  interface{}
}
