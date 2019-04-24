package main

import (
	"learngo/retriever/mock"
	"learngo/retriever/real"
	"fmt"
)
const url = "https://www.imooc.com"

type Retriever interface {
	Get(url string) string
}

func download(r Retriever) string {
	return r.Get(url)
}

type Poster interface {
	Post(url string,form map[string]string) string
}

func post(poster Poster) {
	poster.Post(url,map[string]string{
		"name":"ccmouse",
		"course":"golang",
	})
}

type RetrieverPoster interface {
	Retriever
	Poster
}

func session(s RetrieverPoster) string {
	s.Post(url,map[string]string {
		"contents":"another faked imooc.com",
	})

	return s.Get(url)
}

func main() {
	var r Retriever
	r = &mock.Retriever{Contents:"this is a fake imooc.com"}
	fmt.Println(download(r))

	r = real.Retriever{}
	//fmt.Println(download(r))

	var rp RetrieverPoster
	rp = &mock.Retriever{}
	fmt.Println(session(rp))
}