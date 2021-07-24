package main

import (
	"io/ioutil"
	"os"
)

func main() {
	out, _ := os.Create("boot.go")
	out.Write([]byte("package main\n\nimport (\n"))
	files, _ := ioutil.ReadDir("./backend/")
	BasePkg := "github.com/OrbTools/OrbMap/backend/"
	for _, fil := range files {
		if fil.IsDir() {
			out.Write([]byte("\t_ \"" + BasePkg + fil.Name() + "\"\n"))
		}
	}
	out.Write([]byte(")"))
}
