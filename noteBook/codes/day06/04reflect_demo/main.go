package main

import (
	"fmt"
	"reflect"
)

func reflectType(x interface{}) {
	v := reflect.TypeOf(x)
	fmt.Printf("%v \n", v)
	fmt.Printf("kind=%v name=%v\n", v.Kind(), v.Name())
}

func main() {
	var a int64 = 1024
	reflectType(a)
}
