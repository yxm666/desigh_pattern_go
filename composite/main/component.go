package main

// Component 组件，接口描述了树中简单项目和复杂项目所共有的操作
type Component interface {
	search(keyword string)
}
