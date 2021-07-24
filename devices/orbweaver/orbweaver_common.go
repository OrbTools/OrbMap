package orbweaver

import (
	"fmt"
	"path/filepath"

	morb "github.com/OrbTools/OrbCommon/devices/orbweaver"
	"github.com/OrbTools/OrbMap/registry"
)

type Orbweaver struct {
	eventcodes []byte
	ecm        map[uint16]int
	keymaps    *morb.KeyMaps
}

//ProbcOrbFiles processes orbs
func (p Orbweaver) ProcOrbs(orbs []string) {
	p.keymaps = &morb.KeyMaps{Currentmap: 0}
	if len(orbs) > 0 {
		for idx, orb := range orbs {
			abs, err := filepath.Abs(orb)
			if err != nil {
				panic(err)
			}
			fmt.Println("Loading Orb " + abs)
			KMap := morb.LoadKM(abs)
			p.keymaps.Maps[idx] = KMap
		}
		p.keymaps.MCount = len(orbs)
	} else {
		panic("No orbs")
	}
	p.ecm = make(map[uint16]int)
	p.eventcodes = morb.BINDING[:]
	for i := 0; i < len(p.eventcodes); i++ {
		p.ecm[uint16(p.eventcodes[i])] = i
	}
}

func init() {
	registry.Systems["orbweaver"] = registry.Device(&Orbweaver{})
}
