package main

import (
	"fmt"
	"io"
	"bufio"
	"strings"
	"learngo/functional/fib"
)

type intGen func() int 

//任何类型都能实现接口

//实现Reader接口
func (g intGen) Read(p []byte) (n int,err error) {
	next := g()
	if next > 10000 {
		return 0,io.EOF
	}
	//利用已经实现了Reader的接口 string来做Read
	s := fmt.Sprintf("%d\n",next)

	//incorrect if p is too small!
	return strings.NewReader(s).Read(p)
}

func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	var f intGen = fib.Fibonacci()
	printFileContents(f)

}