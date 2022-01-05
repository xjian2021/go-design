package main

// go run *.go
func main() {
	c := &client{}
	m := &mac{}
	c.insertLightningConnectorIntoComputer(m)

	windowsMachine := &windows{}
	windowsMachineAdapter := &windowsAdapter{
		windowsMachine: windowsMachine,
	}
	c.insertLightningConnectorIntoComputer(windowsMachineAdapter)
}
