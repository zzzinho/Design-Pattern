package main

import "fmt"

func main() {
	client := &Client{}
	mac := &Mac{}

	client.InsertLighteningConnectorIntoComputer(mac)
	fmt.Println()
	windowsMachine := &Windows{}
	windowsMachineAdapter := &WindowsAdapter{
		windowMachine: windowsMachine,
	}

	client.InsertLighteningConnectorIntoComputer(windowsMachineAdapter)
}
