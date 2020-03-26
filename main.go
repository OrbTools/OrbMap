package main

import (
	"flag"
	"os"

	"github.com/minizbot2012/orbmap/orbweaver"
)

func main() {
	var orbs string
	flag.StringVar(&orbs, "orbs", "xiv.orb", "Comma seperated string of orb files")
	flag.Parse()
	path, _ := os.Getwd()
	Maps := orbweaver.ProcOrbFiles(orbs, path)
	orbweaver.OrbLoop(Maps)
}
