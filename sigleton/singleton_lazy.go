package sigleton

import (
	"fmt"
	"sync"
)

/**
懒汉式的单例模式实现,在client主动调用的时候进行对象实例化，注意如果是在多环境下执行要加互斥锁 or sync.Once
*/

var (
	lazySingleton *SingletonStruct

	once = &sync.Once{}

	lock = &sync.Mutex{}
)

func GetInstanceByOncc() *SingletonStruct {
	if singleton == nil {
		once.Do(func() {
			fmt.Println("creating instance")
			lazySingleton = &SingletonStruct{}
		})
	} else {
		fmt.Println("create already")
	}

	return lazySingleton
}

func GetInstanceByLock() *SingletonStruct {
	if singleton == nil {
		// 确保临界区最小，需要临界区的内容尽量都加一个方法，不用影响后续的逻辑在临界区内
		func() {
			lock.Lock()
			defer lock.Unlock()

			if singleton == nil {
				fmt.Println("creating instance")
				lazySingleton = &SingletonStruct{}
			} else {
				fmt.Println("anothercreate already")
			}
		}()

	} else {
		fmt.Println("create instance already")
	}

	return lazySingleton
}
