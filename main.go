package main

import (
	"fmt"
	//"log"
	"gzhcatch/helper"
	"net/url"
)

func main() {
	reqURL := &url.URL{
		Scheme: "http",
		Host:   "www.baidu.com",
		Path:   "/",
	}
	//postdata, _ := url.ParseQuery("")

	var httphelper = &helper.HttpHelper{RequestURL: reqURL}

	bytes, err, status := httphelper.Get()

	if err == nil {
		fmt.Printf("status: %d \r\n", status)
		fmt.Println(len(bytes))
	} else {
		fmt.Println("error:", err.Error(), " status:", status)
		fmt.Println(string(bytes))
	}
}
