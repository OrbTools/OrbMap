//go:generate boxy
//build:+windows

package orbweaver

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"

	"github.com/OrbTools/OrbCommon/devices/common"
	"github.com/OrbTools/OrbCommon/devices/orbweaver"
	"github.com/OrbTools/OrbMap/interface/keyevents"
	"github.com/google/gousb"
)

const (
	vendor = gousb.ID(0x1532)
	prod   = gousb.ID(0x0207)
)

type swaps struct {
	S1 *swapInt
	S2 *swapInt
}

type swapInt struct {
	Modifier byte
	Reserved byte
	K1       byte
	K2       byte
	K3       byte
	K4       byte
	K5       byte
	K6       byte
}

var trans = map[int]int{
	0: 0,
	2: 0,
}

func (s *swapInt) contains(k byte) bool {
	return (s.K1 == k || s.K2 == k || s.K3 == k || s.K4 == k || s.K5 == k || s.K6 == k)
}

func (s *swaps) swap() {
	ss := s.S1
	s.S1 = s.S2
	s.S2 = ss
}
func (s *swapInt) Releases(s2 *swapInt) []byte {
	r := make([]byte, 0)
	if !s.contains(s2.K1) {
		r = append(r, s2.K1)
	}
	if !s.contains(s2.K2) {
		r = append(r, s2.K2)
	}
	if !s.contains(s2.K3) {
		r = append(r, s2.K3)
	}
	if !s.contains((s2.K4)) {
		r = append(r, s2.K4)
	}
	if !s.contains(s2.K5) {
		r = append(r, s2.K5)
	}
	if !s.contains(s2.K6) {
		r = append(r, s2.K6)
	}
	return r
}

func contains(s []byte, e byte) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

//OrbLoop Main loop for this device
func OrbLoop(km *orbweaver.KeyMaps, KeyBus chan *keyevents.KeyEvent) {
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
	swaper := &swaps{}
	swaper.S1 = &swapInt{}
	swaper.S2 = &swapInt{}
	for {
		_, err := rs.Read(data)
		if err != nil {
			panic(err)
		}
		for i := 2; i < in.Desc.MaxPacketSize; i++ {
			if data[i] != 0 {
				data[i] = common.KEYCODE_WINDOWS_FROM_HID[data[i]]
			}
		}
		binary.Read(bytes.NewReader(data), binary.LittleEndian, swaper.S1)
		//data[0] = trans[data[0]]
		for i := 2; i < in.Desc.MaxPacketSize; i++ {
			if data[i] != 0 {
				if common.KEYCODE_WINDOWS_FROM_HID[data[i]] != 255 {
					if !swaper.S2.contains(data[i]) {
						KeyEv := &keyevents.KeyEvent{}
						KeyEv.Code = uint16(data[i])
						KeyEv.Type = 1
						KeyEv.Code = km.Maps[km.Currentmap].Keymap[ecm[KeyEv.Code]]
						KeyBus <- KeyEv
					}
				}
			}
		}
		swaper.swap()
		r := swaper.S2.Releases(swaper.S1)
		for _, rel := range r {
			KeyEv := &keyevents.KeyEvent{}
			KeyEv.Code = uint16(rel)
			KeyEv.Type = 2
			KeyEv.Code = km.Maps[km.Currentmap].Keymap[ecm[KeyEv.Code]]
			KeyBus <- KeyEv
		}
		//Not quite sure how to handle this data quite yet
		println(hex.EncodeToString(data))
	}
}
