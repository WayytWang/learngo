package tree

import (
	"fmt"
)

type Node struct {
	Value int
	Left,Right *Node
}

//结构体的工厂函数
func CreateTreeNode(value int) *Node {
	//返回的是局部变量的地址(c++会报错)
	return &Node{Value:value}
}

//为结构体定义方法
func (node *Node) Print() {
	fmt.Print(node.Value)
}

//结构体的方法的接收者是传值的，需要改变接收者需要传指针
//值接收者 可以用指针来调用
//反之 不行
//nil指针也可以调用方法
func (node *Node) SetValue(value int) {
	if node == nil {
		return 
	}
	node.Value = value
}

func (node *Node) Traverse() {
	if node == nil {
		return
	}

	node.Left.Traverse()
	node.Print()
	node.Right.Traverse()
}

func main() {
	var root Node

	root = Node{Value:3} //只需要设定结构体中多个字段的一个字段默认值，需要key:value
	root.Left = &Node{} 
	root.Right = &Node{5,nil,nil}
	//不论地址还是结构本身，一律使用.来访问成员
	root.Right.Left = new(Node) //开创地址的另一种方法
	root.Left.Right = CreateTreeNode(2)
	root.Right.Left.SetValue(4)

	root.Traverse()
}