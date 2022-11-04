package sigleton

/**
饿汉式的单例模式，会直接实例化对象，而不是等调用时候
*/

// SingletonStruct 模拟要返回的对象
type SingletonStruct struct{}

var singleton *SingletonStruct

// 不对外暴露的初始化方法，会在调用包的时候自动执行init,故会直接执行,全局变量一定有值
func init() {
	singleton = &SingletonStruct{}
}

// GetInstance 对外暴露的获取静态方法的
func GetInstance() *SingletonStruct {
	return singleton
}
