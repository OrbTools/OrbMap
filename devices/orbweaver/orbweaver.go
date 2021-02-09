package orbweaver

import (
	"strings"

	"github.com/OrbTools/OrbCommon/devices/orbweaver"
)

var eventcodes [26]byte
var ecm map[uint16]int

func init() {
	eventcodes = orbweaver.BINDING
	ecm = make(map[uint16]int)
}

//ProcOrbFiles processes orbs
func ProcOrbFiles(orbs string) *orbweaver.KeyMaps {
	keymaps := &orbweaver.KeyMaps{}
	idx := 0
	if len(orbs) > 0 {
		for _, orb := range strings.Split(orbs, ",") {
			KMap := orbweaver.LoadKM(orb)
			keymaps.Maps[idx] = KMap
		}
	} else {
		panic("No orbs")
	}
	return keymaps
}
