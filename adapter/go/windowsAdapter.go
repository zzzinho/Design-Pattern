package main

import "fmt"

type WindowsAdapter struct {
	windowMachine *Windows
}

func (w *WindowsAdapter) InsertIntoLighteningPort() {
	fmt.Println("어댑터로 라이트닝을 USB로 변환")
	w.windowMachine.InsertIntoUSBPort()
}
