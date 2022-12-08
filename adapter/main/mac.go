package main

import "fmt"

// Mac 服务
type Mac struct {
}

func (m *Mac) InsertIntoLightingPort() {
	fmt.Println("Lighting connertor is plugged into mac machine.")
}
