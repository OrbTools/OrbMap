package keypad

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/OrbTools/OrbCommon/devices"
	"github.com/OrbTools/OrbCommon/devices/structs"
	"github.com/OrbTools/OrbMap/registry"
)

type Keypad struct {
	eventcodes []byte
	ecm        map[uint16]int
	keymaps    *structs.KeyMaps
	definition *structs.DeviceDef
}

//ProbcOrbFiles processes orbs
func (p *Keypad) ProcOrbs(dev *structs.DeviceDef, orbs []string) {
	p.definition = dev
	p.keymaps = &structs.KeyMaps{Currentmap: 0}
	if len(orbs) > 0 {
		for _, orb := range orbs {
			abs, _ := filepath.Abs(orb)
			fmt.Println("Loading Orb " + abs)
			file, _ := os.Open(abs)
			KMap := devices.LoadKeymap(file)
			p.keymaps.Maps = append(p.keymaps.Maps, KMap)
		}
		p.keymaps.MCount = len(orbs)
	} else {
		panic("No orbs")
	}
	p.ecm = make(map[uint16]int)
	p.eventcodes = p.definition.Binding
	for i := 0; i < len(p.eventcodes); i++ {
		p.ecm[uint16(p.eventcodes[i])] = i
	}
}

func init() {
	registry.Systems["keypad"] = registry.Device(&Keypad{})
}
