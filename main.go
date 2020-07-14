package main

import (
	"flag"
	"os"

	"github.com/minizbot2012/orbmap/interface/keyevents"
	"github.com/minizbot2012/orbmap/orbweaver"
	"github.com/minizbot2012/orbmap/keypad"
)

func main() {
	var orbs string
	flag.StringVar(&orbs, "orbs", "xiv.orb", "Comma seperated string of orb files")
	flag.Parse()
	path, _ := os.Getwd()
	Maps := orbweaver.ProcOrbFiles(orbs, path)
	KeyBus := make(chan *keyevents.KeyEvent, 128)
	for i := 0; i <= 4; i++ {
		go keypad.ProcKey(KeyBus)
	}
	orbweaver.OrbLoop(Maps, KeyBus)
}
