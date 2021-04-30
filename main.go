//generate:boxy
package main

import (
	"flag"

	"github.com/OrbTools/OrbMap/devices/orbweaver"
	"github.com/OrbTools/OrbMap/emu"
	evdev "github.com/gvalkov/golang-evdev"
)

func main() {
	orbs := flag.String("orbweaver", "", "Comma seperated string of orbs for the orbweaver")
	flag.Parse()
	KeyBus := make(chan *evdev.InputEvent, 128)
	Maps := orbweaver.ProcOrbFiles(*orbs)
	go orbweaver.OrbLoop(Maps, KeyBus)
	emu.ProcKey(KeyBus)
}
