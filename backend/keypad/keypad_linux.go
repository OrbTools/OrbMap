// +build linux !windows
package keypad

import (
	"bytes"
	"encoding/binary"
	"unsafe"

	"github.com/OrbTools/OrbMap/keyevents"
	evdev "github.com/gvalkov/golang-evdev"
)

//OrbLoop Main loop for this device
func (p *Keypad) OrbLoop(KeyBus chan *keyevents.KeyEvent) {
	println("UnixLoop starting")
	f, _ := evdev.Open(p.definition.Device.SystemFile)
	f.Grab()
	var evsize = int(unsafe.Sizeof(keyevents.KeyEvent{}))
	b := make([]byte, evsize)
	for {
		f.File.Read(b)
		KeyEv := new(keyevents.KeyEvent)
		binary.Read(bytes.NewBuffer(b), binary.LittleEndian, KeyEv)
		KeyEv.Code = p.keymaps.Maps[p.keymaps.Currentmap].Keymap[p.ecm[KeyEv.Code]]
		if KeyEv.Code != 0 && KeyEv.Type != 4 {
			KeyBus <- KeyEv
		}
	}
}
