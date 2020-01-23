package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	rsp, err := http.Get("http://localhost:9090/hello?page=12")
	if err != nil {
		fmt.Printf("###############: %v \n", err)
		return
	}
	data, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return
	}
	s := string(data)
	fmt.Printf("[%s]\n", s)
}
