// +build windows

package emu

import (
	"unsafe"

	"github.com/OrbTools/OrbMap/keyevents"
	"github.com/lxn/win"
)

func downKey(key uint16) {
	in := []win.KEYBD_INPUT{
		{
			Type: win.INPUT_KEYBOARD,
			Ki: win.KEYBDINPUT{
				DwExtraInfo: 0,
				WScan:       key,
				DwFlags:     win.KEYEVENTF_SCANCODE,
				Time:        0,
			},
		},
	}
	win.SendInput(1, unsafe.Pointer(&in[0]), int32(unsafe.Sizeof(in[0])))
}
func upKey(key uint16) {
	in := []win.KEYBD_INPUT{
		{
			Type: win.INPUT_KEYBOARD,
			Ki: win.KEYBDINPUT{
				DwExtraInfo: 0,
				WScan:       key,
				DwFlags:     win.KEYEVENTF_KEYUP | win.KEYEVENTF_SCANCODE,
				Time:        0,
			},
		},
	}
	win.SendInput(1, unsafe.Pointer(&in[0]), int32(unsafe.Sizeof(in[0])))
}

//ProcKey Windows support is so weird
func ProcKey(kb chan *keyevents.KeyEvent) {
	println("Emu Windows Starting")
	for {
		KeyEv := <-kb
		if KeyEv.Type == 1 {
			downKey(KeyEv.Code)
		} else if KeyEv.Type == 2 {
			upKey(KeyEv.Code)
		}
	}
}
