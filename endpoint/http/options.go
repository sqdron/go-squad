package http

import (
	"github.com/sqdron/go-squad/optionizer"
)

func (t *HttpTransport) Options() []interface{} {
	return optionizer.Group("Http Endpoint Configuration",
		optionizer.NewOption("addr").Short("a").Param("host").Description("Bind to host address"),
		optionizer.NewOption("port").Short("p").Param("port").Description("Use port for clients"))
}
