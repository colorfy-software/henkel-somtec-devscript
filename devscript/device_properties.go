package devscript

import (
	"github.com/sirupsen/logrus"
	"reflect"
)

type DeviceProperties struct {
	Properties        map[string]*Property
	OnPropertyChanged func(prop *Property)
}

type Property struct {
	Name      string
	Writeable bool
	Value     any
	Kind      reflect.Kind
}

func (dp *DeviceProperties) Set(name string, value any) {
	if prop, ok := dp.Properties[name]; ok {
		if reflect.TypeOf(value).Kind() != prop.Kind {
			panic("invalid type for property " + name)
		}
		if prop.Value != value {
			prop.Value = value
			dp.OnPropertyChanged(prop)
			logrus.Infof("'%s' set to '%v'", name, value)
		}
		return
	}
	panic("property not found: " + name)
}

func (dp *DeviceProperties) Get(name string) any {
	if prop, ok := dp.Properties[name]; ok {
		return prop.Value
	}
	panic("property not found: " + name)
}
