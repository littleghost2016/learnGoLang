package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	// fmt.Println("Hello, world@")
	// url := "http://www.baidu.com"
	// download(url)
}

func download(url string) {
	client := &http.Client{}
	requset, _ := http.NewRequest("GET", url, nil)

	//自定义Header
	requset.Header.Set("User-Agent", "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1)")

	response, err := client.Do(requset)
	if err != nil {
		fmt.Println("Http get error", err)
		return
	}

	//函数结束后关闭相关链接
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Read error", err)
		return
	}
	fmt.Println(string(body))
}