package main

import (
	"fmt"
	"net/http"
)

func main() {
	resp, err := http.Get("https://www.bilibili.com/bangumi/play/ss36168")
	if err != nil {
		fmt.Println("http.Get err: ", err)
		return
	}
	defer resp.Body.Close()
	//fmt.Println("Status=", resp.Status)
	//fmt.Println("StatusCode=", resp.StatusCode)
	//fmt.Println("Header=", resp.Header)
	//fmt.Println("Body=", resp.Body)
	fmt.Println(resp)
}
