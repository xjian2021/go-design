package main

import "fmt"

type client struct {}

func (c *client) insertLightningConnectorIntoComputer(com computer)  {
	fmt.Println("客户端将Lightning连接器插入计算机")
	com.insertIntoLightningPort()
}
