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
		//处理未能预计到的错误
		defer func() {
			if r := recover(); r!= nil {
				log.Printf("Panic: %v",r)
				http.Error(writer,http.StatusText(http.StatusInternalServerError),http.StatusInternalServerError)
			}
		}()

		err := handler(writer,request)
		if err != nil {
			log.Printf("Error handling request: %s",err.Error())

			//判断错误是不是userError
			if userErr,ok := err.(userError);ok {
				http.Error(writer,userErr.Message(),http.StatusBadRequest)
				return 
			}

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

type userError interface {
	//继承error接口
	error
	Message() string
}

func main() {
	http.HandleFunc("/",errWrapper(filelisting.HanldeFileList))

	err := http.ListenAndServe(":8888",nil) 
	if err != nil {
		panic(err)
	}
}