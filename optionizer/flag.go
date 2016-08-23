package optionizer

import (
	"flag"
	"fmt"
	"os"
)

func (op *optionizer) ParseFlags() {
	for _, v := range op.options {
		option:= v.(*Option);
		flag.Var(option, option.name, option.description)
		if (option.short != ""){
			flag.Var(option, option.short, option.description)
		}
	}
	flag.Usage = func() {
		fmt.Println(op.String())
		os.Exit(0)
	}
	flag.Parse()
}