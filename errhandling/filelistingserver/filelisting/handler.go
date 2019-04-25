package filelisting

import (
	"net/http"
	"os"
	"io/ioutil"
	"strings"
)

const prefix = "/list/"

//实现了main中userError接口
type userError string

//实现error接口
func (e userError) Error() string {
	return e.Message()
}

func (e userError) Message() string {
	return string(e)
}

func HanldeFileList(writer http.ResponseWriter,request *http.Request) error {
	if strings.Index(request.URL.Path,prefix) != 0 {
		return userError("path must start with" + prefix)
	}
	
	path := request.URL.Path[len(prefix):]

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

