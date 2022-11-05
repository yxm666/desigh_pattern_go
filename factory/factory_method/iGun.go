package factory_method

/**
由于Go中缺少类和继承等OOP特性，无法使用Go来实现经典的工厂方法模式。不过，我们仍能实现模式的基础版本，即简单工厂。
*/

// 创建了一个名为IGun的接口，其中将定义一支枪所需具备的所有方法
type IGun interface {
	setName(name string)
	setPower(power int)
	getName() string
	getPower() int
}
