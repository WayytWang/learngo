package filelisting

import (
	"net/http"
	"os"
	"io/ioutil"
)

func HanldeFileList(writer http.ResponseWriter,request *http.Request) error {
	path := request.URL.Path[len("/list/"):]

	file,err := os.Open(path)
	if err != nil {
		//http返回错误
		//http.Error(wirter,err.Error(),http.StatusInternalServerError)
		return err
	}
	defer file.Close()

	all,err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	writer.Write(all)
	return nil
}

