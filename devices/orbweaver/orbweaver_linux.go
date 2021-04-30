// +build linux !windows

package orbweaver

import (
	"bytes"
	"encoding/binary"

	morb "github.com/OrbTools/OrbCommon/devices/orbweaver"
	evdev "github.com/gvalkov/golang-evdev"
)

//OrbLoop Main loop for this device
func OrbLoop(km *morb.KeyMaps, KeyBus chan *evdev.InputEvent) {
	eventcodes = morb.BINDING[:]
	for i := 0; i < len(eventcodes); i++ {
		ecm[uint16(eventcodes[i])] = i
	}
	println("UnixLoop starting")
	f, err := evdev.Open("/dev/input/by-id/usb-Razer_Razer_Orbweaver_Chroma-event-kbd")
	if err != nil {
		panic(err)
	}
	f.Grab()
	b := make([]byte, 24)
	for {
		KeyEv, _ := f.ReadOne()
		binary.Read(bytes.NewReader(b[16:]), binary.LittleEndian, &KeyEv)
		KeyEv.Code = km.Maps[km.Currentmap].Keymap[ecm[KeyEv.Code]]
		if KeyEv.Code != 0 && KeyEv.Type != 4 {
			KeyBus <- KeyEv
		}
	}
}
