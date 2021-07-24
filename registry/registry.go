package registry

import (
	"github.com/OrbTools/OrbMap/keyevents"
)

var (
	Systems map[string]Device
)

type Device interface {
	OrbLoop(chan *keyevents.KeyEvent)
	ProcOrbs([]string)
}

func init() {
	Systems = make(map[string]Device)
}
