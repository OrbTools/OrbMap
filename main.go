//generate:boxy
package main

import (
	"flag"
	"os"
	"strings"

	"github.com/minizbot2012/orbmap/box"
	"github.com/minizbot2012/orbmap/devices/orbweaver"
	"github.com/minizbot2012/orbmap/interface/keyevents"
	"github.com/minizbot2012/orbmap/keypad"
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
	keypad.ProcKey(KeyBus)
}
