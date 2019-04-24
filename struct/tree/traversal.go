package tree

import (
	"fmt"
)
//用函数式编程来实现遍历

//原来的遍历函数
// func (node *Node) Traverse() {
// 	if node == nil {
// 		return
// 	}

// 	node.Left.Traverse()
// 	node.Print()
// 	node.Right.Traverse()
// }


func (node *Node) Traverse() {
	node.TraverseFunc(func(node *Node){
		node.Print()
	})
	fmt.Println()
}

//改造
func (node *Node) TraverseFunc(f func(*Node)) {
	if node == nil {
		return
	}

	node.Left.TraverseFunc(f)
	f(node)
	node.Right.TraverseFunc(f)
}
