//go:generate boxy

package orbweaver

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/bendahl/uinput"
)

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
func OrbLoop(km *KeyMaps) {
	var EventCodes = box.Get("orbweaver.dev")
	ecm := make(map[uint16]int)
	for i := 0; i < 26; i++ {
		ecm[EventCodes[i]] = i
	}
	f, err := os.Open("/dev/input/by-id/usb-Razer_Razer_Orbweaver_Chroma-event-kbd")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	b := make([]byte, 24)
	vkm, _ := uinput.CreateKeyboard("/dev/uinput", []byte("Orbmap"))
	defer vkm.Close()
	for {
		f.Read(b)
		sec := binary.LittleEndian.Uint64(b[0:8])
		usec := binary.LittleEndian.Uint64(b[8:16])
		t := time.Unix(int64(sec), int64(usec))
		var _ = t
		var value int32
		typ := binary.LittleEndian.Uint16(b[16:18])
		code := binary.LittleEndian.Uint16(b[18:20])
		binary.Read(bytes.NewReader(b[20:]), binary.LittleEndian, &value)
		/*fmt.Println(t)
		fmt.Println(typ)
		fmt.Println(code)
		fmt.Println(value)*/
		if typ == 1 {
			if value == 1 {
				vkm.KeyDown(km.Maps[km.Currentmap].Keymap[ecm[code]])
			} else if value == 2 {
				//pass
			} else {
				vkm.KeyUp(km.Maps[km.Currentmap].Keymap[ecm[code]])
			}
		}
	}
}
