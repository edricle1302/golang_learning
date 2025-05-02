package adapter

import (
	"fmt"
	"testing"
)

func TestAdapter(t *testing.T) {
	// main.go
	fmt.Println("*** Example Adapter ***")

	client := &Client{}
	mac := &Mac{}

	client.InsertLightningConnectorIntoComputer(mac)

	windowsMachine := &Windows{}
	windowsMachineAdapter := &WindowsAdapter{
		windowMachine: windowsMachine,
	}

	client.InsertLightningConnectorIntoComputer(windowsMachineAdapter)

	fmt.Print("*** End of Adapter ***\n\n\n")
}
