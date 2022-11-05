package factory_method

import (
	"fmt"
	"testing"
)

func TestOne(t *testing.T) {
	ak47, _ := getGun("ak47")
	musket, _ := getGun("musket")

	printDetails(ak47)
	printDetails(musket)
}

func printDetails(g IGun) {
	// 调用具体产品的方法
	fmt.Printf("Gun: %s", g.getName())
	fmt.Println()
	// 调用具体产品的方法
	fmt.Printf("Power: %d", g.getPower())
	fmt.Println()
}
