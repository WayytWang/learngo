package main

import (
	"fmt"
)

type treeNode struct {
	value int
	left,right *treeNode
}

//结构体的工厂函数
func createTreeNode(value int) *treeNode {
	//返回的是局部变量的地址(c++会报错)
	return &treeNode{value:value}
}

//为结构体定义方法
func (node *treeNode) print() {
	fmt.Print(node.value)
}

//结构体的方法的接收者是传值的，需要改变接收者需要传指针
//值接收者 可以用指针来调用
//反之 不行
//nil指针也可以调用方法
func (node *treeNode) setValue(value int) {
	if node == nil {
		return 
	}
	node.value = value
}

func (node *treeNode) traverse() {
	if node == nil {
		return
	}

	node.left.traverse()
	node.print()
	node.right.traverse()
}

func main() {
	var root treeNode

	root = treeNode{value:3} //只需要设定结构体中多个字段的一个字段默认值，需要key:value
	root.left = &treeNode{} 
	root.right = &treeNode{5,nil,nil}
	//不论地址还是结构本身，一律使用.来访问成员
	root.right.left = new(treeNode) //开创地址的另一种方法
	root.left.right = createTreeNode(2)
	root.right.left.setValue(4)

	root.traverse()
}