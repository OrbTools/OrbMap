//go:generate boxy

package orbweaver

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
	"strings"

	"github.com/minizbot2012/orbmap/box"
	"github.com/minizbot2012/orbmap/interface/keyevents"
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

//OrbLoop Main loop for this device
func OrbLoop(km *KeyMaps, KeyBus chan keyevents.KeyEvent) {
	for i := 0; i < 26; i++ {
		ecm[uint16(eventcodes[i])] = i
	}
	f, err := os.Open("/dev/input/by-id/usb-Razer_Razer_Orbweaver_Chroma-event-kbd")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	b := make([]byte, 24)
	for {
		f.Read(b)
		KeyEv := keyevents.KeyEvent{}
		KeyEv.Type = binary.LittleEndian.Uint16(b[16:18])
		KeyEv.Code = km.Maps[km.Currentmap].Keymap[ecm[binary.LittleEndian.Uint16(b[18:20])]]
		binary.Read(bytes.NewReader(b[20:]), binary.LittleEndian, &KeyEv.Value)
		if KeyEv.Code != 0 {
			KeyBus <- KeyEv
		}
	}
}
