package main

import (
	"flag"
	"os"

	"github.com/bendahl/uinput"
	"github.com/minizbot2012/orbmap/interface/keyevents"
	"github.com/minizbot2012/orbmap/orbweaver"
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

func main() {
	var orbs string
	flag.StringVar(&orbs, "orbs", "xiv.orb", "Comma seperated string of orb files")
	flag.Parse()
	path, _ := os.Getwd()
	Maps := orbweaver.ProcOrbFiles(orbs, path)
	KeyBus := make(chan keyevents.KeyEvent, 128)
	go procKey(KeyBus)
	orbweaver.OrbLoop(Maps, KeyBus)
}
