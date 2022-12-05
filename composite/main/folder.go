package main

import "fmt"

type Folder struct {
	// 含有复杂项目和简单项目
	components []Component
	name       string
}

func (f *Folder) search(keyword string) {
	fmt.Printf("Searching recursively for keyword %s in folder %s\n", keyword, f.name)
	// 遍历复杂成员的所有子成员，调用不同实现了共有接口(component)的方法进行遍历
	// 将所有工作委派给子元素
	for _, composite := range f.components {
		composite.search(keyword)
	}
}

func (f *Folder) add(c Component) {
	f.components = append(f.components, c)
}
