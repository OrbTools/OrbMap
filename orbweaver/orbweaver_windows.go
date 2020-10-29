package orbweaver

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/minizbot2012/orbmap/interface/keyevents"
	"github.com/minizbot2012/winorb/sys/keyboard"
	"github.com/minizbot2012/winorb/sys/types"
)

//OrbLoop Main loop for this device
func OrbLoop(km *KeyMaps, KeyBus chan *keyevents.KeyEvent) {
	for i := 0; i < 26; i++ {
		ecm[uint16(eventcodes[i])] = i
	}
	keyboardChan := make(chan types.KeyboardEvent, 16)
	if err := keyboard.Install(nil, keyboardChan); err != nil {
		panic(err)
	}

	defer keyboard.Uninstall()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)

	fmt.Println("start capturing keyboard input")

	for {
		k := <-keyboardChan
		fmt.Printf("Received %V %v\n", k.Message, k.ScanCode)
		KeyEv := &keyevents.KeyEvent{}
		if k.Flags == 0x2 {
			KeyEv.Type = 2
		} else {
			KeyEv.Type = 1
		}
		KeyEv.Code = uint16(k.ScanCode)
		KeyEv.Code = km.Maps[km.Currentmap].Keymap[ecm[KeyEv.Code]]
	}
}
