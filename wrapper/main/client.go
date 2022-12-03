package main

import "fmt"

/*
	客户端代码，用于接受一个对象(ligntning接口)的部分功能，还有一个名为adaptee的对象(windows笔记本)，可以通过不同的接口(USB接口)实现相同的功能

	创建一个adapter的结构体:
	1. 遵循符合客户端期望的相同的接口
	2. 可以适合被适配对象的方式倆对客户端的请求进行"翻译"。适配器能够接受来自Lighting连接器的信息，并将其转换成USB格式的信号，同时将信号传递给Windows笔记本的USB接口
*/

type Client struct {
}

func (c *Client) InsertLightningConnectorIntoComputer(com Computer) {
	fmt.Println("Client inserts Lighting connerctor into computer.")
	com.InsertIntoLightingPort()
}
