package main

import (
	"learngo/struct/queue"
	"fmt"
)

func main() {
	q := queue.Queue{1}

	//切片类型本身就是指针
	q.Push(2)
	q.Push(3)
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())
}