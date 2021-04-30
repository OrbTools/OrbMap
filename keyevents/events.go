package keyevents

import "syscall"

//KeyEvent represents a keyevent
type KeyEvent struct {
	Time  syscall.Timeval
	Type  uint16
	Code  uint16
	Value int32
}
