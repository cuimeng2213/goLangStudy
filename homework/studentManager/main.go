package main

import "os"
import "fmt"

// 学生管理系统 使用结构体和方法实现
type Student struct {
	name string
	age  int
}
type StudentManager struct {
	allBooks []*Student
}

func newStudent(name string, age int) *Student {
	return &Student{
		name: name,
		age:  age,
	}
}

func (s *StudentManager) AddStudent() {
	var (
		name string
		age  int
	)
	fmt.Println("g根据提示信息操作")
	fmt.Print("学生名字：")
	fmt.Scan(&name)
	fmt.Print("学生年龄:")
	fmt.Scan(&age)
	stu := newStudent(name, age)

	s.allBooks = append(s.allBooks, stu)

}
func (s *StudentManager) deleteStudent() {
	var name string
	flag := false
	fmt.Print("请输入要删除的学生姓名：")
	fmt.Scan(&name)

	for index, stu := range s.allBooks {
		if stu.name == name {
			// 删除
			s.allBooks = append(s.allBooks[:index], s.allBooks[index+1:]...)
			fmt.Printf("%s 同学删除成功", name)
			flag = true
		}
	}
	if !flag {
		fmt.Printf("无此 %s 同学", name)
	}

}
func (s *StudentManager) updateStudent() {
	var (
		name string
		age  int
	)
	fmt.Println("g根据提示信息操作")
	fmt.Print("学生名字：")
	fmt.Scan(&name)
	// 查看是否有该学生
	for _, stu := range s.allBooks {
		if stu.name == name {
			fmt.Print("修改学生名字:")
			fmt.Scan(&name)
			fmt.Print("修改学生年龄:")
			fmt.Scan(&age)
			//stu := newStudent(name, age)
			stu.name = name
			stu.age = age
			fmt.Println("修改成功")
			return
		}
	}

	fmt.Printf("查无此%s 学生", name)

}
func (s *StudentManager) ShowAllStudents() {
	for _, s := range s.allBooks {
		fmt.Printf("名字：%s 年龄：%d \n", s.name, s.age)
	}
}

func show_help() {
	fmt.Println("welcome Student Manager System!")
	fmt.Println("1: 添加学生")
	fmt.Println("2: 删除学生")
	fmt.Println("3: 更新学生信息")
	fmt.Println("4: 显示所有学生")
	fmt.Println("5: 退出")
}

func main() {
	var option int
	manager := &StudentManager{}
	manager.allBooks = make([]*Student, 0, 200)
	for {
		show_help()
		fmt.Scan(&option)
		switch option {
		case 1:
			manager.AddStudent()
		case 2:
			manager.deleteStudent()
		case 3:
			manager.updateStudent()
		case 4:
			manager.ShowAllStudents()
		case 5:
			os.Exit(0)
		}
	}
}
