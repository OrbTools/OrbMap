//generate:boxy
package main

import (
	"flag"
	"strings"

	"github.com/OrbTools/OrbMap/emu"
	"github.com/OrbTools/OrbMap/keyevents"
	"github.com/OrbTools/OrbMap/registry"
)

//go:generate go run generators/devices.go
func main() {
	str := make(map[string]*string)
	for d := range registry.Systems {
		str[d] = flag.String(d, "", "Comma seperated list of orb files for "+d)
	}
	flag.Parse()
	KeyBus := make(chan *keyevents.KeyEvent, 128)
	for sys, orbs := range str {
		if len(*orbs) > 0 {
			registry.Systems[sys].ProcOrbs(strings.Split(*orbs, ","))
			go registry.Systems[sys].OrbLoop(KeyBus)
		}
	}
	emu.ProcKey(KeyBus)
}
