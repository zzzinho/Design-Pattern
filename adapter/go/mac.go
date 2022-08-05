package main

import "fmt"

type Mac struct {
}

func (m *Mac) InsertIntoLighteningPort() {
	fmt.Println("라이트닝 케이블이 맥에 플러그되었다.")
}
