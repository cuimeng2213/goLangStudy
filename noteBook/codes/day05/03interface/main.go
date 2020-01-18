//03_interface
package main
import (
	"fmt"
)
//引出接口实例

type speaker interface{
	speak()
}

type cat struct{}
type dog struct{}
type person struct{}

func (c *cat)speak(){
	fmt.Println("miao")
}
func (c *dog)speak(){
	fmt.Println("wang")
}
func (c *person)speak(){
	fmt.Println("aaaaa")
}

//
// func da(x person){
// 	//接收一个参数，传进什么来，我就打什么
// 	x.speak()
// }
//
func da(x speaker){
	//接收一个参数，传进什么来，我就打什么
	x.speak()
}

func main(){
	p := &person{}
	da(p)
}