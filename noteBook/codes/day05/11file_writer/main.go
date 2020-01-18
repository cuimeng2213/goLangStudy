package main

import (
	"fmt"
	"os"
)

func demo() {

}

func main() {
	fp, err := os.OpenFile("./xx.txt", os.O_APPEND|os.O_CREATE, 0644)
	defer fp.Close()
	if err != nil {
		fmt.Printf("Open file error: %s\n", err)
		return
	}
	fmt.Printf("%T\n", fp)

	n, err := fp.Write([]byte("hello"))
	if err != nil {
		fmt.Printf("write file Error")
		return
	}
	fmt.Printf("write %d\n", n)
}
