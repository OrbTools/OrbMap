package orbweaver

import (
	"encoding/hex"
	"math/big"

	"github.com/google/gousb"
	"github.com/minizbot2012/orbmap/devices/common"
	"github.com/minizbot2012/orbmap/interface/keyevents"
)

const (
	vendor = gousb.ID(0x1532)
	prod   = gousb.ID(0x0207)
)

//OrbLoop Main loop for this device
func OrbLoop(km *KeyMaps, KeyBus chan *keyevents.KeyEvent) {
	for i := 0; i < 26; i++ {
		ecm[uint16(eventcodes[i])] = i
	}
	ctx := gousb.NewContext()
	dev, err := ctx.OpenDeviceWithVIDPID(vendor, prod)
	if err != nil {
		panic(err)
	}
	defer dev.Close()
	conf, err := dev.Config(1)
	if err != nil {
		panic(err)
	}
	intf, err := conf.Interface(0, 0)
	if err != nil {
		panic(err)
	}
	defer intf.Close()
	in, err := intf.InEndpoint(1)
	if err != nil {
		panic(err)
	}
	data := make([]byte, in.Desc.MaxPacketSize)
	rs, _ := in.NewStream(in.Desc.MaxPacketSize, 3)
	var bits big.Int
	for {
		_, err := rs.Read(data)
		if err != nil {
			panic(err)
		}
		for i := 2; i < in.Desc.MaxPacketSize; i++ {
			if common.KEYCODE_WINDOWS_FROM_HID[data[i]] == 255 {
				data[i] = 0
			} else {
				data[i] = common.KEYCODE_WINDOWS_FROM_HID[data[i]]
			}
			if data[i] != 0 && bits.Bit(ecm[uint16(data[i])]) == 0 {
				KeyEv := &keyevents.KeyEvent{
					Code: uint16(data[i]),
					Type: 1,
				}
				bits.SetBit()
				KeyEv.Code = km.Maps[km.Currentmap].Keymap[ecm[KeyEv.Code]]
			} else if
		}
		//Not quite sure how to handle this data quite yet
		println(hex.EncodeToString(data))
	}
}
