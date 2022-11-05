package factory_method

// 具体产品
type Ak47 struct {
	Gun
}

// 工厂子方法，返回一个接口
func newAk47() IGun {
	return &Ak47{Gun: Gun{name: "ak47", power: 4}}
}
