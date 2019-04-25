package main

import (
	"net/http"
	"os"
	"log"
	"learngo/errhandling/filelistingserver/filelisting"
)

type appHandler func(writer http.ResponseWriter,request *http.Request) error

func errWrapper(handler appHandler) func(http.ResponseWriter,*http.Request){
	return func(writer http.ResponseWriter,request *http.Request) {
		err := handler(writer,request)
		if err != nil {
			log.Printf("Error handling request: %s",err.Error())
			code := http.StatusOK
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			default:
				code = http.StatusInternalServerError
			}
			http.Error(writer,http.StatusText(code),code)
		}
	}
}

func main() {
	http.HandleFunc("/list/",errWrapper(filelisting.HanldeFileList))

	err := http.ListenAndServe(":8888",nil) 
	if err != nil {
		panic(err)
	}
}