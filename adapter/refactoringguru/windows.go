package main

import "fmt"

type windows struct {}

func (w *windows) insertIntoUSBPort() {
	fmt.Println("USB接口已插入windows计算机")
}
