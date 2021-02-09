//build:+windows

package emu

import (
	"syscall"
	"unsafe"

	"github.com/OrbTools/OrbMap/interface/keyevents"
	"github.com/lxn/win"
)

var user32 = syscall.NewLazyDLL("user32.dll")
var procKeyBd = user32.NewProc("keybd_event")
var mapVK = user32.NewProc("MapVirtualKeyA")

func downKey(key int) {
	flag := 0
	var in win.KEYBD_INPUT
	vs, _, _ := mapVK.Call(uintptr(uint32(key)), uintptr(1))
	vsc := uint16(vs)
	println(vsc)
	if vsc == win.VK_SHIFT || vsc == win.VK_CONTROL || vsc == win.VK_MENU {
		flag |= win.KEYEVENTF_EXTENDEDKEY
	}
	in.Type = 1
	in.Ki.DwExtraInfo = 0
	in.Ki.DwFlags = uint32(flag)
	in.Ki.WScan = 0
	in.Ki.WVk = vsc

	win.SendInput(1, unsafe.Pointer(&in), int32(unsafe.Sizeof(in)))
}
func upKey(key int) {
	flag := win.KEYEVENTF_KEYUP
	var in win.KEYBD_INPUT
	vs, _, _ := mapVK.Call(uintptr(uint32(key)), uintptr(1))
	vsc := uint16(vs)
	if vsc == win.VK_SHIFT || vsc == win.VK_CONTROL || vsc == win.VK_MENU {
		flag |= win.KEYEVENTF_EXTENDEDKEY
	}
	in.Type = 1
	in.Ki.DwExtraInfo = 0
	in.Ki.DwFlags = uint32(flag)
	in.Ki.WScan = 0
	in.Ki.WVk = vsc
	win.SendInput(1, unsafe.Pointer(&in), int32(unsafe.Sizeof(in)))
}

//ProcKey Windows support is so weird
func ProcKey(kb chan *keyevents.KeyEvent) {
	for {
		KeyEv := <-kb
		if KeyEv.Type == 1 {
			downKey(int(KeyEv.Code))
		} else if KeyEv.Type == 2 {
			upKey(int(KeyEv.Code))
		}
	}
}
