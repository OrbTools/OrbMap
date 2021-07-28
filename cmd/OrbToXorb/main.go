//OrbToXorb will convert old orb files to the new Xorb format
//XOrb will contain new features eventually, however this
//Is an interium measure while converting to xdr2 within
//OrbMap, OrbBind and OrbCommon
package main

import (
	"encoding/binary"
	"flag"
	"io"
	"os"
	"strings"

	"github.com/OrbTools/OrbCommon/devices"
	xdr "github.com/davecgh/go-xdr/xdr2"
)

func loadOrb(file string, dev *devices.DeviceDef) *devices.KeyMap {
	mapped := new(devices.KeyMap)
	of, _ := os.Open(file)
	defer of.Close()
	mapped.Keymap = make([]uint16, dev.NumKeys)
	binary.Read(of, binary.LittleEndian, mapped.Keymap)
	binary.Read(of, binary.LittleEndian, mapped.Color)
	return mapped
}
func writeXorb(mapped interface{}, file io.WriteCloser) {
	xdr.Marshal(file, mapped)
	file.Close()
}
func main() {
	devt := flag.String("dev", "", "Device type to convert")
	inorb := flag.String("orb", "", "orb file to convert to xorb")
	flag.Parse()
	km := loadOrb(*inorb, devices.DeviceTypes[*devt])
	file, _ := os.Create(strings.Split(*inorb, ".")[0] + ".xorb")
	writeXorb(km, file)
}
