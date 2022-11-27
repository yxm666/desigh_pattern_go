package main

import (
	"fmt"
)

/*
抽象工厂接口
*/
type ISportsFactory interface {
	makeShoe() IShoe
	makeShirt() IShirt
}

// GetSportsFactory 获取具体的工厂的方法
func GetSportsFactory(brand string) (ISportsFactory, error) {
	if brand == "adidas" {
		return &Adidas{}, nil
	}

	if brand == "nike" {
		return &Nike{}, nil
	}

	return nil, fmt.Errorf("Wrong brand type passed")
}
