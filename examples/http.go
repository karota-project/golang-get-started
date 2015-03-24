package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func byteToString(c []byte) string {
	n := -1
	for i, b := range c {
		if b == 0 {
			break
		}
		n = i
	}
	return string(c[:n+1])
}

func main() {
	res, err := http.Get("http://www.google.co.jp")
	if err != nil {
		fmt.Println(fmt.Println(err))
	}
	fmt.Println("http-header : ", res)

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(fmt.Println(err))
	}
	fmt.Println("http-body(byte) : ", body)

	str := byteToString(body[:])
	fmt.Println("http-body(string) : ", str)
}
