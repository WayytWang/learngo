package main

import (
	"fmt"
	"sync"
)

//使用sync.WaitGroup来等待全部协程执行完毕

func doWorker(id int,c chan int,w worker) {
	for n := range c {
		fmt.Printf("Worker %d received %c \n", id, n)
		w.done
	}
}

//wg *sync.WaitGroup

type worker struct {
	in chan int
	done func() 
}

func createWorker(id int,wg *sync.WaitGroup) worker {
	w := worker {
		in : make(chan int),
		done : func(){
			wg.Done()
		},
	}
	go doWorker(id,w.in,wg)
	return w 
}

func chanDemo() {
	//要保证所有的协程中都是同一个wg
	var wg sync.WaitGroup
	var workers [10]worker

	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i,&wg)
	}

	wg.Add(20)
	for i,worker := range workers {
		worker.in <- 'a' + i
	}

	for i,worker := range workers {
		worker.in <- 'A' + i
	}
	wg.Wait()

}


func main() {
	chanDemo()
}