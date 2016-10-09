package configurator

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"reflect"
	"strings"
	"log"
)

type Configurator struct {
}

func New() *Configurator {
	return &Configurator{}
}

func (cfg *Configurator) mapOptions(options interface{}) {
	val := reflect.ValueOf(options).Elem()
	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i)
		typeField := val.Type().Field(i)
		name := strings.ToLower(typeField.Name)
		description := typeField.Tag.Get("option")
		valuePointer := valueField.Addr().Interface()
		switch valueField.Kind() {
		case reflect.String:
			flag.StringVar(valuePointer.(*string), name, "", description)
		case reflect.Int:
			flag.IntVar(valuePointer.(*int), name, 0, description)
		}
	}
}

func (cfg *Configurator) ReadFromFile(file string, options interface{}) {

	files, _ := ioutil.ReadDir("./")
	for _, f := range files {
		fmt.Println(f.Name())
		log.Println(f.Name())
	}

	data, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = json.Unmarshal(data, &options)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func (cfg *Configurator) ReadFlags(options interface{}) {
	config := ""
	flag.StringVar(&config, "config", "", "Config file path")
	cfg.mapOptions(options)
	flag.Parse()
	if config != "" {
		cfg.ReadFromFile(config, options)
	}
}
