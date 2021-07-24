// +build linux !windows

package orbweaver

import (
	"bytes"
	"encoding/binary"
	"unsafe"

	"github.com/OrbTools/OrbMap/keyevents"
	evdev "github.com/gvalkov/golang-evdev"
)

//OrbLoop Main loop for this device
func (p *Orbweaver) OrbLoop(KeyBus chan *keyevents.KeyEvent) {
	println("UnixLoop starting")
	f, err := evdev.Open("/dev/input/by-id/usb-Razer_Razer_Orbweaver_Chroma-event-kbd")
	if err != nil {
		panic(err)
	}
	f.Grab()
	var evsize = int(unsafe.Sizeof(keyevents.KeyEvent{}))
	b := make([]byte, evsize)
	for {
		f.File.Read(b)
		KeyEv := &keyevents.KeyEvent{}
		binary.Read(bytes.NewBuffer(b), binary.LittleEndian, KeyEv)
		KeyEv.Code = p.keymaps.Maps[p.keymaps.Currentmap].Keymap[p.ecm[KeyEv.Code]]
		if KeyEv.Code != 0 && KeyEv.Type != 4 {
			KeyBus <- KeyEv
		}
	}
}
