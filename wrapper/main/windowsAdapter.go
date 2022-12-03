package main

import "fmt"

// windowsAdapter 适配器
type WindowsAdapter struct {
	windowMachine *Windows
}

// client接口直接调用适配器，而不去调用服务
func (w *WindowsAdapter) InsertIntoLightingPort() {
	fmt.Println("Adapter converts Lighting signal to USB")
	w.windowMachine.insertIntoUSBPort()
}
