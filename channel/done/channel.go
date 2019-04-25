package main

import (
	"fmt"
)

//done管道来做协程之间的通信
func doWorker(id int,c chan int,done chan bool) {
	for n := range c {
		fmt.Printf("Worker %d received %c \n", id, n)
		done <- true
	}
}
 
type worker struct {
	in chan int
	done chan bool
}

func createWorker(id int) worker {
	w := worker {
		in : make(chan int),
		done : make(chan bool),
	}
	go doWorker(id,w.in,w.done)
	return w 
}

func chanDemo() {
	var workers [10]worker
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i)
	}
	
	for i,worker := range workers {
		worker.in <- 'a' + i
		//确保读完之后再写
		//<- workers[i].done
	}

	//等待所有写入in的数据被读出来之后在结束程序
	for _,worker := range workers {
		<-worker.done
	}

	for i,worker := range workers {
		worker.in <- 'A' + i
		//<- worker.done
	}

	//等待所有写入in的数据被读出来之后在结束程序
	for _,worker := range workers {
		<-worker.done
	}
}


func main() {
	chanDemo()
}