package orbweaver

import (
	"encoding/binary"
	"fmt"
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
	Keymap [26]int
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
	keymaps := new(KeyMaps)
	idx := 0
	fmt.Println(wd + ":" + orbs)
	if len(orbs) > 0 {
		for _, orb := range strings.Split(orbs, ",") {
			KMap := new(KeyMap)
			inf, _ := os.Open(wd + "/" + orb)
			for i := 0; i < 26; i++ {
				b := make([]byte, 2)
				inf.Read(b)
				KMap.Keymap[i] = int(binary.LittleEndian.Uint16(b))
			}
			keymaps.Maps[idx] = KMap
			idx++
			inf.Close()
		}
	} else {
		panic("No orbs")
	}
	return keymaps
}
