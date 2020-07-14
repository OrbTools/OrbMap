package keypad

import (
	"github.com/bendahl/uinput"
	"github.com/minizbot2012/orbmap/interface/keyevents"
)

var vkm uinput.Keyboard = nil

func init() {
	var err error
	vkm, err = uinput.CreateKeyboard("/dev/uinput", []byte("Orbmap"))
	if err != nil {
		panic(err)
	}
}

//ProcKey keyboard emulator loop
func ProcKey(kb chan *keyevents.KeyEvent) {
	for vkm == nil {
		println("VKM not init")
	}
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
