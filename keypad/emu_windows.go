package keypad

import (
	"unsafe"

	"github.com/lxn/win"
	"github.com/minizbot2012/orbmap/interface/keyevents"
)

type keyboardInput struct {
	wVk         uint16
	wScan       uint16
	dwFlags     uint32
	time        uint32
	dwExtraInfo uint64
}

type input struct {
	inputType uint32
	ki        keyboardInput
	padding   uint64
}

//Windows support is so weird
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
