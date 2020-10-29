package orbweaver

import (
	"github.com/minizbot2012/orbmap/interface/keyevents"
)

//OrbLoop Main loop for this device
func OrbLoop(km *KeyMaps, KeyBus chan *keyevents.KeyEvent) {
	// TODO: NEED HELP!!
	for i := 0; i < 26; i++ {
		ecm[uint16(eventcodes[i])] = i
	}
}
