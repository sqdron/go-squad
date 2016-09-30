package http
//
//import (
//	"github.com/sqdron/squad/optionizer"
//)
//
//type Options struct {
//	Host string
//	Port string
//}
//
//func (t *HttpTransport) Configuratoin() []interface{} {
//	return optionizer.Group("Http Endpoint Configuration",
//		optionizer.NewOption(&t.Options.Host, "addr").Short("a").Param("host").Description("Bind to host address"),
//		optionizer.NewOption(&t.Options.Port, "port").Short("p").Param("port").Description("Use port for clients"))
//}
