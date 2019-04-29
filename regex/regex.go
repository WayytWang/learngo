package main

import (
	"fmt"
	"regexp"
)

const text = `My email is ccmouse@gmail.com
wyj@gmail.com
www@gmail.com.cn`

// .代表任何字符
// +代表一个或者多个 *代表0个或者多个
// ``中的内容不包含转义字符
// 正则表达式中() 用来提取()中的内容

func main() {
	re := regexp.MustCompile(`([a-zA-Z0-9]+)@([a-zA-Z0-9]+)(\.[a-zA-Z0-9.]+)`)
	//match := re.FindString(text)  //提取一个
	//match := re.FindAllString(text, -1) //提取多个
	match := re.FindAllStringSubmatch(text, -1)
	for _, m := range match {
		fmt.Println(m)
	}

}
