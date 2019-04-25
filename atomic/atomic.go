package main

//atomic有线程安全的加法 这里自己实现一下

import (
	//"sync/atomic"
	"fmt"
	"time"
	"sync"
)

type atomicInt struct{
	value int 
	lock sync.Mutex
}

func (a *atomicInt) increment() {
	a.lock.Lock()
	defer a.lock.Unlock()
	a.value++
}

func (a *atomicInt) get() int {
	a.lock.Lock()
	defer a.lock.Unlock()
	return a.value
}

func main() {
	var a atomicInt
	a.increment()
	go func() {
		a.increment()
	}()
	time.Sleep(time.Millisecond)
	fmt.Println(a.get())
}