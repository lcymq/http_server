package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	resp, err := http.Get("http://127.0.0.1:9090/posts/Go/15_socket/")
	if err != nil {
		fmt.Println("get url failed, err: %v\n", err)
		return
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("read resp.Body failed, err:", err)
	}
	fmt.Println(string(b))
}
