package main

import (
	"learngo/functional/fib"
	"os"
	"bufio"
	"fmt"
)

func tyrDefer() {
	for i := 0; i < 100; i++ {
		defer fmt.Println(i) //参数在defer语句时计算
		if i == 30 {
			panic("printed too many")
		}
	}
}

func writeFile(filename string) {
	file,err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	f := fib.Fibonacci()
	for i := 0; i < 20 ;i++ {
		fmt.Fprintln(writer, f())
	}
}

func main() {
	tyrDefer()
	writeFile("fib.txt")
}