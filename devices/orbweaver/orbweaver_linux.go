//go:generate boxy

package orbweaver

import (
	"bytes"
	"encoding/binary"
	"os"

	"github.com/minizbot2012/orbmap/interface/keyevents"
)

//OrbLoop Main loop for this device
func OrbLoop(km *KeyMaps, KeyBus chan *keyevents.KeyEvent) {
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
		KeyEv := &keyevents.KeyEvent{}
		binary.Read(bytes.NewReader(b[16:]), binary.LittleEndian, &KeyEv)
		KeyEv.Code = km.Maps[km.Currentmap].Keymap[ecm[KeyEv.Code]]
		if KeyEv.Code != 0 && KeyEv.Type != 4 {
			KeyBus <- KeyEv
		}
	}
}
