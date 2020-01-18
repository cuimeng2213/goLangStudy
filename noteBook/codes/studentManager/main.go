package main

import (
	"fmt"
)

type Student struct{
	id int64
	name string
}

func showMenu(){
	fmt.Println("Welecome student manager system")
	fmt.Println(`
		1. 添加学生
		2. 显示学生
		3. 删除学生
		4. 退出
	`)
}
var stuMgr *stuManager

func main(){
	var cmd int
	stuMgr = NewStuManager()
	for{
		showMenu()
		fmt.Print("Please choice a num: ")
		_, err := fmt.Scan(&cmd)
		if err != nil {
			fmt.Println("input error")
			return 
		}
		switch cmd{
		case 1:
			stuMgr.AddStudent()
		case 2:
			stuMgr.ShowStudent()
		case 3:
			stuMgr.DelStudent()
		case 4:
			return
		}
	}
}