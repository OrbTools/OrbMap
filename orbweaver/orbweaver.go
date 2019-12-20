package orbweaver

import (
	"strings"
	"encoding/binary"
	"os"
	"time"
	"fmt"
	"bytes"
	"github.com/bendahl/uinput"
)

type KeyMap struct {
	Keymap [26]int
	Color [3]byte
}
type KeyMaps struct {
	Maps [7]*KeyMap
	Currentmap int 
	MCount int
}
func Proc_orb_files(orbs string, wd string) *KeyMaps {
	keymaps := new(KeyMaps)
	idx := 0
	fmt.Println(wd + ":" +orbs)
	if(len(orbs) > 0) {
		for _, orb := range strings.Split(orbs, ",") {
			KMap := new(KeyMap)
			inf, _ := os.Open(wd+"/"+orb)
			for i := 0; i<26; i++ {
				b := make([]byte, 2)
				inf.Read(b);
				KMap.Keymap[i] = int(binary.LittleEndian.Uint16(b))
			}
			keymaps.Maps[idx] = KMap
			idx+=1;
			inf.Close()
		}
	} else {
		panic("No orbs")
	}
	return keymaps
}
func OrbLoop(km *KeyMaps) {
	var event_codes = [...]uint16{41, 2, 3, 4, 5, 15, 16, 17, 18, 19, 58, 30, 31, 32, 33, 42, 44, 45, 46, 47, 56, 103, 106, 108, 105, 57}
	ecm := make(map[uint16]int)
	for i := 0; i<26; i++ {
		ecm[event_codes[i]] = i
	}
	fmt.Println(string(ev))
	f, err := os.Open("/dev/input/by-id/usb-Razer_Razer_Orbweaver_Chroma-event-kbd")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	b := make([]byte, 24)
	vkm, _ := uinput.CreateKeyboard("/dev/uinput", []byte("Orbmap"))
	defer vkm.Close()
	for {
		f.Read(b)
		sec := binary.LittleEndian.Uint64(b[0:8])
		usec := binary.LittleEndian.Uint64(b[8:16])
		t := time.Unix(int64(sec), int64(usec))
		var _ = t
		var value int32
		typ := binary.LittleEndian.Uint16(b[16:18])
		code := binary.LittleEndian.Uint16(b[18:20])
		binary.Read(bytes.NewReader(b[20:]), binary.LittleEndian, &value)
		/*fmt.Println(t)
		fmt.Println(typ)
		fmt.Println(code)
		fmt.Println(value)*/
		if(typ == 1) {
			if(value == 1 || value == 2) {
				vkm.KeyDown(km.Maps[km.Currentmap].Keymap[ecm[code]])
			} else {
				vkm.KeyUp(km.Maps[km.Currentmap].Keymap[ecm[code]])
			}
		}
	}
}