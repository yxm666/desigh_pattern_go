package main

import "fmt"

// Mac 精确抽象
type Mac struct {
	printer Printer
}

func (m *Mac) Print() {
	fmt.Println("Print request for mac")
	m.printer.PrintFile()
}

// SetPrinter 装载实现
func (m *Mac) SetPrinter(p Printer) {
	m.printer = p
}
