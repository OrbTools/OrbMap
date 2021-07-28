package registry

import (
	"reflect"

	"github.com/OrbTools/OrbCommon/devices"
	"github.com/OrbTools/OrbMap/keyevents"
)

var (
	Systems map[string]Device
)

type Device interface {
	OrbLoop(chan *keyevents.KeyEvent)
	ProcOrbs(*devices.DeviceDef, []string)
}

func init() {
	Systems = make(map[string]Device)
}

func NewOf(name string) Device {
	nInter := reflect.New(reflect.ValueOf(Systems[name]).Type().Elem())
	return nInter.Interface().(Device)
}
