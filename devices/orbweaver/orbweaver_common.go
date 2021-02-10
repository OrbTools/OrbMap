package orbweaver

import (
	"fmt"
	"path/filepath"
	"strings"

	morb "github.com/OrbTools/OrbCommon/devices/orbweaver"
)

var eventcodes []byte
var ecm map[uint16]int

func init() {
	ecm = make(map[uint16]int)
}

//ProcOrbFiles processes orbs
func ProcOrbFiles(orbs string) *morb.KeyMaps {
	keymaps := &morb.KeyMaps{Currentmap: 0}
	if len(orbs) > 0 {
		for idx, orb := range strings.Split(orbs, ";") {
			abs, err := filepath.Abs(orb)
			if err != nil {
				panic(err)
			}
			fmt.Println("Loading Orb " + abs)
			KMap := morb.LoadKM(abs)
			keymaps.Maps[idx] = KMap
		}
		keymaps.MCount = len(orbs)
	} else {
		panic("No orbs")
	}
	return keymaps
}
