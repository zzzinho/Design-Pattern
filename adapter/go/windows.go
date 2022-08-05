package main

import "fmt"

type Windows struct {
}

func (w *Windows) InsertIntoUSBPort() {
	fmt.Println("USB가 윈도우에 플러그 되었다.")
}
