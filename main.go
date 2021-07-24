//generate:boxy
package main

import (
	"flag"
	"strings"

	"github.com/OrbTools/OrbCommon/devices"
	"github.com/OrbTools/OrbMap/emu"
	"github.com/OrbTools/OrbMap/keyevents"
	"github.com/OrbTools/OrbMap/registry"
)

//go:generate go run generators/backends.go
func main() {
	str := make(map[string]*string)
	for d, dev := range devices.DeviceTypes {
		str[d] = flag.String(d, "", "Comma seperated list of orb files for "+d+" "+dev.Backend)
	}
	flag.Parse()
	KeyBus := make(chan *keyevents.KeyEvent, 128)
	for sys, orbs := range str {
		if len(*orbs) > 0 {
			devh := registry.NewOf(devices.DeviceTypes[sys].Backend)
			devh.ProcOrbs(devices.DeviceTypes[sys], strings.Split(*orbs, ","))
			go devh.OrbLoop(KeyBus)
		}
	}
	emu.ProcKey(KeyBus)
}
