// +build windows

package orbweaver

import (
	"bytes"
	"encoding/binary"
	"fmt"

	morb "github.com/OrbTools/OrbCommon/devices/orbweaver"
	"github.com/OrbTools/OrbCommon/hid"
	"github.com/OrbTools/OrbMap/keyevents"
	"github.com/google/gousb"
)

const (
	vendor           = gousb.ID(0x1532)
	prod             = gousb.ID(0x0207)
	leftControl byte = 0x1
	leftShift   byte = 0x2
	leftAlt     byte = 0x4
)

type swaps struct {
	S1 *swapInt
	S2 *swapInt
}

type swapInt struct {
	K1 byte
	K2 byte
	K3 byte
	K4 byte
	K5 byte
	K6 byte
	K7 byte
	K8 byte
	K9 byte
}

func (s *swapInt) contains(k byte) bool {
	return (s.K1 == k || s.K2 == k || s.K3 == k || s.K4 == k || s.K5 == k || s.K6 == k || s.K7 == k || s.K8 == k || s.K9 == k)
}

func (s *swaps) swap() {
	ss := s.S1
	s.S1 = s.S2
	s.S2 = ss
}

func trans(M byte) []byte {
	r := make([]byte, 0)
	if (M & leftShift) != 0 {
		r = append(r, byte(hid.GetMappingFromName("SHIFT_LEFT").Evdev))
	} else {
		r = append(r, 0)
	}
	if (M & leftControl) != 0 {
		r = append(r, byte(hid.GetMappingFromName("CONTROL_LEFT").Evdev))
	} else {
		r = append(r, 0)
	}
	if (M & leftAlt) != 0 {
		r = append(r, byte(hid.GetMappingFromName("SHIFT_ALT").Evdev))
	} else {
		r = append(r, 0)
	}
	return r
}

func (s *swapInt) Differ(s2 *swapInt) []byte {
	r := make([]byte, 0)
	if !s2.contains(s.K1) {
		r = append(r, s.K1)
	}
	if !s2.contains(s.K2) {
		r = append(r, s.K2)
	}
	if !s2.contains(s.K3) {
		r = append(r, s.K3)
	}
	if !s2.contains(s.K4) {
		r = append(r, s.K4)
	}
	if !s2.contains(s.K5) {
		r = append(r, s.K5)
	}
	if !s2.contains(s.K6) {
		r = append(r, s.K6)
	}
	if !s2.contains(s.K7) {
		r = append(r, s.K7)
	}
	if !s2.contains(s.K8) {
		r = append(r, s.K8)
	}
	if !s2.contains(s.K9) {
		r = append(r, s.K9)
	}
	return r
}

//OrbLoop Main loop for this device
func OrbLoop(km *morb.KeyMaps, KeyBus chan *keyevents.KeyEvent) {
	eventcodes = morb.BINDING[:]
	for i := 0; i < len(eventcodes); i++ {
		ecm[uint16(eventcodes[i])] = i
	}
	fmt.Println("Windows Loop Init")
	ctx := gousb.NewContext()
	dev, err := ctx.OpenDeviceWithVIDPID(vendor, prod)
	if err != nil {
		panic(err)
	}
	fmt.Println("Device connected")
	defer dev.Close()
	conf, err := dev.Config(1)
	if err != nil {
		panic(err)
	}
	intf, err := conf.Interface(0, 0)
	if err != nil {
		panic(err)
	}
	fmt.Println("Windows Loop Interf")
	defer intf.Close()
	in, err := intf.InEndpoint(1)
	if err != nil {
		panic(err)
	}
	fmt.Println("Windows Loop Pointed")
	data := make([]byte, in.Desc.MaxPacketSize)
	rs, _ := in.NewStream(in.Desc.MaxPacketSize, 3)
	swaper := new(swaps)
	swaper.S1 = new(swapInt)
	swaper.S2 = new(swapInt)
	fmt.Println("Windows Loop Starting")
	for {
		_, err := rs.Read(data)
		if err != nil {
			panic(err)
		}
		addin := trans(data[0])
		tdat := data[2:]
		dat := append(addin, tdat...)
		for i := 0; i < len(dat); i++ {
			if dat[i] != 0 {
				dat[i] = byte(hid.GetLinuxFromHid(uint16(dat[i])))
				dat[i] = byte(km.Maps[km.Currentmap].Keymap[ecm[uint16(dat[i])]])
				dat[i] = byte(hid.GetHidFromLinux(uint16(dat[i])))
				dat[i] = byte(hid.GetWindowsFromHid(uint16(dat[i])))
			}
		}
		err = binary.Read(bytes.NewReader(dat), binary.LittleEndian, swaper.S1)
		if err != nil {
			panic(err)
		}
		for _, pre := range swaper.S1.Differ(swaper.S2) {
			if pre != 0 {
				KeyEv := &keyevents.KeyEvent{}
				KeyEv.Code = uint16(pre)
				KeyEv.Type = 1
				KeyBus <- KeyEv
			}
		}
		for _, rel := range swaper.S2.Differ(swaper.S1) {
			if rel != 0 {
				KeyEv := &keyevents.KeyEvent{}
				KeyEv.Code = uint16(rel)
				KeyEv.Type = 2
				KeyBus <- KeyEv
			}
		}
		swaper.swap()
	}
}
