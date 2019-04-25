package main

import (
	"fmt"
	"time"
)

//main函数自身也是goroutine

func main() {
	var a [10]int
	for i := 0;i < 10;i++ {
		go func(i int) {
			for {
				//如果i直接引用循环中的i 很不安全
				//fmt.Printf()是io操作，会切换协程 (这里的切换不是协程本身切换的，协程本身是非抢占式的)
				//fmt.Printf("Hello from goroutine %d \n", i)
				a[i]++
			}
		}(i)
	}

	time.Sleep(time.Millisecond)
	fmt.Println(a)
}

				