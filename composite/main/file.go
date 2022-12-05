/*
操作系统的文件系统的例子，文件系统中有两种类型的对象:文件和文件夹。在某些情形中，文件和文件夹应被视为相同的对象。这就是组合模式发挥作用的时候了
*/
package main

import "fmt"

// File 组件接口 叶子节点
type File struct {
	name string
}

// 实现了公共类接口的叶子节点
func (f *File) search(keyword string) {
	fmt.Printf("Searching for keyword %s in file %s\n", keyword, f.name)
}

func (f *File) getName() string {
	return f.name
}
