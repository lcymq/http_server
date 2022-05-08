package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	resp, err := http.Get("http://127.0.0.1:9090/testhttp/")
	if err != nil {
		fmt.Println("get url failed, err: %v\n", err)
		return
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("read resp.Body failed, err:", err)
	}
	fmt.Println(string(b))

	data := url.Values{}
	urlObj, _ := url.Parse("http://127.0.0.1:9000/xxx/")
	data.Set("name", "moqi")
	data.Set("age", "16")
	queryStr := data.Encode()
	fmt.Println(queryStr)
	urlObj.RawQuery = queryStr
	req, err := http.NewRequest("GET", urlObj.String(), nil)
	resp, err = http.DefaultClient.Do(req)
}
