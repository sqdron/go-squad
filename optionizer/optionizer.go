package optionizer

import (
	"github.com/sqdron/go-squad/query"
	"fmt"
)

type Option struct {
	name        string
	params      []string
	group       string
	description string
	short       string
	value       interface{}
}

func (o *Option) String() string {
	if (o.value != nil) {
		return o.value.(string)
	}
	return ""

}
func (o *Option) Set(v string) error {
	o.value = v
	return nil
}

type optionizer struct {
	options []interface{}
}

type ApplyFunc func(op *optionizer)


func New() *optionizer {
	return &optionizer{};
}

func (o *optionizer) Range(options []interface{}) *optionizer {
	for _, v := range options {
		o.options = append(o.options, v)
	}
	return o;
}

func (o *optionizer) Group(name string, options ...OptionFunc) *optionizer {
	o.Range(Group(name, options...))
	return o;
}

func Group(name string, options ...OptionFunc) []interface{} {
	res:= []interface{}{}
	for _, v := range options {
		option := &Option{}
		option.group = name
		v(option)
		res = append(res, option)
	}
	return res;
}

type OptionFunc func(o *Option)

func NewOption(name string) OptionFunc {
	return func(o *Option) {
		o.name = name
	}
}

func (of OptionFunc) Short(name string) OptionFunc {
	return func(o *Option) {
		of(o)
		o.short = name
	}
}

func (of OptionFunc) Param(name string) OptionFunc {
	return func(o *Option) {
		of(o)
		o.params = append(o.params, name)
	}
}

func (of OptionFunc)Description(description string) OptionFunc {
	return func(o *Option) {
		of(o)
		o.description = description
	}
}

func (of OptionFunc) Default(value interface{}) OptionFunc {
	return func(o *Option) {
		of(o)
		o.value = value
	}
}

func (op *optionizer) String() string {
	grouped := query.GroupBy(op.options, func(x interface{}) interface{} {
		return x.(*Option).group;
	});
	result := "\n";
	for group, options := range grouped {
		result += fmt.Sprintf("%s:\n", group)
		for _, option := range options {
			o := option.(*Option)
			pms := "	"
			if (o.short != "") {
				pms += "-" + o.short + ","
			}
			pms += "--" + o.name;

			if (len(o.params) > 0) {
				pms += " <"
				for i, v := range o.params {

					if (i != 0) {
						pms += ", "
					}
					pms += v
				}
				pms += ">"
			}

			result += fmt.Sprintf("%-30s %s", pms, o.description)

			result += "\n"
		}
	}
	return result;
}

func (op *optionizer) Print() {

}

//type ApplyFunc func(*Option) *optionizer
//type OptionFunc func() ApplyFunc
//
//func (op *optionizer) Group(name string) ApplyFunc {
//	return func(o *Option) *optionizer {
//		if (o != nil) {
//			o.group = name
//			op.options = append(op.options, o)
//		}
//		return op;
//	}
//}
//
//func (a ApplyFunc) Option(name, short, description string) OptionFunc {
//	return func() ApplyFunc {
//		option := &Option{name:name, description:description, short:short}
//		a(option)
//		return a
//	}
//}
//
//func (of OptionFunc) Option(name, short, description string) OptionFunc {
//	return func() ApplyFunc {
//		a := of();
//		a.Option(name, short, description)()
//		return a;
//	}
//}
////
////func (of OptionFunc) Default(df interface{}) OptionFunc {
////	return func() ApplyFunc {
////		a := of();
////		a.Option(name, short, description)()
////		return a;
////	}
////}
//
//func (op OptionFunc) Group(name string) ApplyFunc {
//	return op.Done().Group(name)
//}
//
//func (g OptionFunc) Done() *optionizer {
//	return g()(nil)
//}
//
//func (o *optionizer) Option(name, short, description string) *optionizer {
//	o.options = append(o.options, &Option{name:name, description:description, short:short})
//	return o
//}
//
