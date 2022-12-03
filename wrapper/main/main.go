package main

func main() {
	client := &Client{}
	mac := &Mac{}
	// 直接调用
	client.InsertLightningConnectorIntoComputer(mac)

	// 通过适配器调用
	windowsMachine := &Windows{}
	windowsMachineAdapter := &WindowsAdapter{
		windowMachine: windowsMachine,
	}
	client.InsertLightningConnectorIntoComputer(windowsMachineAdapter)
}
