// +build linux !windows

package emu

import (
	"github.com/bendahl/uinput"
	evdev "github.com/gvalkov/golang-evdev"
)

var vkm uinput.Keyboard = nil

//ProcKey keyboard emulator loop
func ProcKey(kb chan *evdev.InputEvent) {
	println("Emu Unix Starting")
	var err error
	vkm, err = uinput.CreateKeyboard("/dev/uinput", []byte("Orbmap"))
	if err != nil {
		panic(err)
	}
	defer vkm.Close()
	for {
		KeyEv := <-kb
		if KeyEv.Code == 1 {
			if KeyEv.Value == 1 {
				vkm.KeyDown(int(KeyEv.Code))
			} else if KeyEv.Value == 2 {
			} else {
				vkm.KeyUp(int(KeyEv.Code))
			}
		}
	}
}
