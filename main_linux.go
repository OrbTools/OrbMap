package main

import (
	"github.com/bendahl/uinput"
	"github.com/minizbot2012/orbmap/interface/keyevents"
)

func procKey(kb chan keyevents.KeyEvent) {
	vkm, _ := uinput.CreateKeyboard("/dev/uinput", []byte("Orbmap"))
	defer vkm.Close()
	for {
		KeyEv := <-kb
		if KeyEv.Type == 1 {
			if KeyEv.Value == 1 {
				vkm.KeyDown(KeyEv.Code)
			} else if KeyEv.Value == 2 {
			} else {
				vkm.KeyUp(KeyEv.Code)
			}
		}
	}
}
