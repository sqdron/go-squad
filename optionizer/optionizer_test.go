package optionizer

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"github.com/smartystreets/assertions/should"
	"fmt"
	"flag"
	"os"
)

func Test_Group_Optionizer(t *testing.T) {
	o := New().Group("Server Options:",
		NewOption("addr").Short("a").Param("host").Description("Bind to host address").Default("0.0.0.0"),
		NewOption("port").Short("p").Param("port").Description("Use port for clients").Default("3333"),
		NewOption("debug").Description("Debug and trace")).
	Group("TLS Options:",
		NewOption("tls").Description("Enable TLS, do not verify clients (default: false)"))

	Convey("Make composition", t, func() {
		So(o, should.NotBeNil)
		fmt.Println(o)
	})

	flag.Usage = func(){
		fmt.Println(o);
		os.Exit(0)
	}
	o.ParseFlags()
}
