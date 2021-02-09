//generate:boxy
package main

import (
	"flag"
	"os"
	"strings"

	"github.com/OrbTools/OrbMap/box"
	"github.com/OrbTools/OrbMap/devices/orbweaver"
	"github.com/OrbTools/OrbMap/emu"
	"github.com/OrbTools/OrbMap/interface/keyevents"
)

func main() {
	var orbs []*string
	for _, v := range box.List() {
		println(v)
		orbs = append(orbs, flag.String(strings.Split(v, ".")[0], "", "Comma seperated string of orbs for the orbweaver"))
	}
	flag.Parse()
	path, _ := os.Getwd()
	KeyBus := make(chan *keyevents.KeyEvent, 128)
	if *orbs[0] != "" {
		Maps := orbweaver.ProcOrbFiles(*orbs[0], path)
		go orbweaver.OrbLoop(Maps, KeyBus)
	}
	emu.ProcKey(KeyBus)
}
