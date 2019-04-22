package main 

import (
	"learngo/struct/tree"
	"fmt"
)

//扩充Node结构体
type myTreeNode struct {
	node *tree.Node
}

//后序遍历
func (myNode *myTreeNode) postOrder() {
	if myNode == nil || myNode.node == nil {
		return
	}
	left := myTreeNode{myNode.node.Left}
	left.postOrder()

	right := myTreeNode{myNode.node.Right}
	right.postOrder()
	myNode.node.Print()

}


func main() {
	var root tree.Node

	root = tree.Node{Value:3} //只需要设定结构体中多个字段的一个字段默认值，需要key:value
	root.Left = &tree.Node{} 
	root.Right = &tree.Node{Value:5,Left:nil,Right:nil}
	//不论地址还是结构本身，一律使用.来访问成员
	root.Right.Left = new(tree.Node) //开创地址的另一种方法
	root.Left.Right = tree.CreateTreeNode(2)
	root.Right.Left.SetValue(4)

	root.Traverse()
	fmt.Println()
	myroot := myTreeNode{&root}
	myroot.postOrder()
	fmt.Println("")
}