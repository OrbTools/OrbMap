package main

import (
	"io/ioutil"
	"os"
)

func main() {
	out, _ := os.Create("boot.go")
	out.Write([]byte("package main\n\nimport (\n"))
	files, _ := ioutil.ReadDir("./devices/")
	BasePkg := "github.com/OrbTools/OrbMap/devices/"
	for _, fil := range files {
		if fil.IsDir() {
			out.Write([]byte("\t_ \"" + BasePkg + fil.Name() + "\"\n"))
		}
	}
	out.Write([]byte(")"))
}
