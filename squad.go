package squad

// squad is an isolated micro-service unit. Nobody knows about his existence but it knows about the entire world.
type service struct {
	id      string
	Version string
	Name    string
	//config   *configurator
	//endpoint endpoint.IEndpoint
	//transport transport.ITransport

}

type IService interface {
	Activate()
}

func New() *service {
	return &service{}
}

//
//func (s *squad) Transport(t transport.ITransport) *squad{
//	s.transport = t;
//}

//Activate and forgot
func (s *service) Activate() {

}
