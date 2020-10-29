package keypad

import (
	"unsafe"

	"github.com/lxn/win"
	"github.com/minizbot2012/orbmap/interface/keyevents"
)

//ProcKey Windows support is so weird
func ProcKey(kb chan *keyevents.KeyEvent) {
	for {
		KeyEv := <-kb
		var i win.KEYBD_INPUT
		i.Type = 1
		i.Ki.WScan = uint16(KeyEv.Code)
		i.Ki.DwFlags = 0x8
		if KeyEv.Type == 2 {
			i.Ki.DwFlags |= 0x2
		}
		win.SendInput(1, unsafe.Pointer(&i), int32(unsafe.Sizeof(i)))
	}
}
