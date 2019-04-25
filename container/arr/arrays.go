package main

import (
	"fmt"
)

func main() {
	var arr1 [5]int //声明
	arr2 := [3]int{1,3,5} //声明并复制
	arr3 := [...]int{2,4,6,8,10}
	var grid [4][5]bool

	fmt.Println(arr1,arr2,arr3,grid)
}