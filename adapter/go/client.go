package main

import "fmt"

type Client struct {
}

func (c *Client) InsertLighteningConnectorIntoComputer(com Computer) {
	fmt.Println("클라이언트가 라이트닝 케이블을 컴퓨터에 플러그했다.")
	com.InsertIntoLighteningPort()
}
