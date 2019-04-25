package main

import (
	"net/http"
	"fmt"
	"net/http/httputil"
)

func main() {
	request,err := http.NewRequest(http.MethodGet,"http://www.imooc.com",nil)
	request.Header.Add("User-Agent",
		"Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Mobile/15A372 Safari/604.1")

	client := http.Client{
		//req 这一次的重定向 via：所有的重定向 error=nil允许重定向
		CheckRedirect : func(req *http.Request,via []*http.Request) error {
			fmt.Println("Redirect:",req)
			return nil
		},
	}
	resp,err := client.Do(request)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	//转储 (获得resp.Body里面的内容)
	s,err := httputil.DumpResponse(resp,true)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", s)
}