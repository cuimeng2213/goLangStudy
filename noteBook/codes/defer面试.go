package main

import "fmt"

func fn() int {
	x := 5
	defer func() {
		x++
		fmt.Println("inner: ", x)
	}()
	fmt.Println("after inner: ", x)
	return x
}

func fn2() (x int) {
	defer func() {
		x++
		fmt.Println("inner: ", x)
	}()
	fmt.Println("after inner: ", x)
	return x
}

func main() {
	fmt.Println(fn())

	fmt.Println(fn2())
}
