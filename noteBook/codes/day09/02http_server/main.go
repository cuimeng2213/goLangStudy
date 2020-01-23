package main

import (
	"fmt"
	"net/http"
)

func f1(w http.ResponseWriter, r *http.Request) {
	//http的get请求方法，参数是放在url中的？问号后面。
	//http://www.baidu.com?search=baidu&page=12
	param := r.URL.Query()
	fmt.Printf("%s = %s \n", "page", param["page"])
	fmt.Fprintf(w, "ok")
}
func main() {
	http.HandleFunc("/hello", f1)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Printf("AAAAAAAAA err: %v\n", err)
		return
	}
}
