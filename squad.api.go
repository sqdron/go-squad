package squad

type ApiFunc func()

func (s *squad) Api() ApiFunc

type ApiAction func(request interface{}) interface{}

func (af ApiFunc) Action(path string, action ApiAction){

}


