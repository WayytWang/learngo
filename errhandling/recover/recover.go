package main

import (
	//"errors"
	"fmt"
)

func tryRecover() {
	defer func() {
		// r 是panic出来的东西
		r := recover()
		if err,ok := r.(error);ok {
			fmt.Println("Error occurred:",err)
		}else {
			panic(r)
		}
	}()

	//panic(interface{})
	//panic(errors.New("this is an error"))
	b := 0
	a := 5/b
	fmt.Println(a)
}

func main() {
	tryRecover()
}