package main

import (
	"fmt"
	"time"
)

func worker(id int, c chan int) {
	//可以用range的方式来遍历管道，管道关闭之后会跳出
	// for n := range c {

	// }

	for {
		//判断channel是否关闭了
		n, ok := <-c
		//ok == false表示管道已经关闭了，关闭时n是管道类型的零值
		if !ok {
			break
		}
		fmt.Printf("Worker %d received %c \n", id, n)
	}
}

func createWorker(id int) chan int {
	c := make(chan int)
	go func() {
		for {
			fmt.Printf("Worker %d received %c \n", id, <-c)
		}
	}()
	return c
}

func chanDemo() {
	var channels [10]chan int
	for i := 0; i < 10; i++ {
		// channels[i] = make(chan int)
		// go worker(i,channels[i])

		channels[i] = createWorker(i)
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i
	}
	time.Sleep(time.Millisecond)
}

func bufferedChannel() {
	c := make(chan int, 3) //可以接收三个(在没有取出的情况下)
	go worker(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	time.Sleep(time.Millisecond)
}

//channel的关闭是由发送方关闭，来告诉结束方写数据完毕
func channleClose() {
	c := make(chan int)
	go worker(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	close(c)
	time.Sleep(time.Millisecond)
}

func main() {
	fmt.Println("Channel as first-class citizen")
	chanDemo()

	fmt.Println("Buffered channel")
	bufferedChannel()

	fmt.Println("Channel close and range")
	channleClose()
}
