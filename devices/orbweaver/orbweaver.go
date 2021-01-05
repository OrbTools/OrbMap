package orbweaver

import (
	"encoding/binary"
	"os"
	"strings"

	"github.com/minizbot2012/orbmap/box"
)

var eventcodes []byte
var ecm map[uint16]int

func init() {
	eventcodes = box.Get("orbweaver.dev")
	ecm = make(map[uint16]int)
}

//KeyMap singular keymap
type KeyMap struct {
	Keymap [26]uint16
	Color  [3]byte
}

//KeyMaps a set of keymaps
type KeyMaps struct {
	Maps       [7]*KeyMap
	Currentmap int
	MCount     int
}

//ProcOrbFiles processes orbs
func ProcOrbFiles(orbs string, wd string) *KeyMaps {
	keymaps := &KeyMaps{}
	idx := 0
	if len(orbs) > 0 {
		for _, orb := range strings.Split(orbs, ",") {
			KMap := &KeyMap{}
			inf, _ := os.Open(wd + "/" + orb)
			defer inf.Close()
			err := binary.Read(inf, binary.LittleEndian, KMap)
			if err != nil {
				panic(err)
			}
			keymaps.Maps[idx] = KMap
		}
	} else {
		panic("No orbs")
	}
	return keymaps
}
