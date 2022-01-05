package main

import "fmt"

type windowsAdapter struct {
	windowsMachine *windows
}

func (w *windowsAdapter) insertIntoLightningPort() {
	fmt.Println("适配器将lightning信号转换为USB")
	w.windowsMachine.insertIntoUSBPort()
}


