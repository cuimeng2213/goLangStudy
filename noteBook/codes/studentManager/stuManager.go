package main
import(
	"fmt"
)
//学生管理结构体
type stuManager struct{
	allStudents map[int64]*Student
}

// 新增学生
func (m *stuManager) AddStudent(){
	var(
		id int64
		name string
	)
	fmt.Print("请输入学生学号：")
	fmt.Scan(&id)
	fmt.Print("请输入学生名字：")
	fmt.Scan(&name)
	m.allStudents[id] = &Student{
		id:id,
		name:name,
	}
}

func (m *stuManager) ShowStudent(){
	fmt.Println("=======================")
	for _, stu := range m.allStudents{
		fmt.Printf("学号: %d 姓名: %s \n",stu.id, stu.name)
	}
	fmt.Println("=======================")
}
func (m *stuManager) DelStudent(){
	var(
		id int64
		//name string
	)
	fmt.Print("请输入学生学号：")
	fmt.Scan(&id)

	if !m.hasId(id){
		return
	}

	delete(m.allStudents,id)
}

func (m *stuManager) ModifyStudent(){
	var(
		id int64
		//name string
	)
	fmt.Print("请输入学生学号：")
	fmt.Scan(&id)
	//先查找对应id的学生是否存在
	if !m.hasId(id) {
		return
	}
}

func (m *stuManager) hasId(id int64) bool{
	_, ok := m.allStudents[id]
	if !ok {
		fmt.Printf("No have this %d id student! ", id)
		return false
	}
	return true
}

func NewStuManager() *stuManager{
	return &stuManager{
		allStudents : make(map[int64]*Student, 100),
	}
}