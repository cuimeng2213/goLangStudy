package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"reflect"
	"strconv"
	"strings"
)

type MysqlConfig struct {
	Address  string `ini:"address"`
	Port     int    `ini:"port"`
	Username string `ini:"username"`
	Password string `ini:"password"`
}
type RedisConfig struct {
	Host     string `ini:"host"`
	Port     int    `ini:"port"`
	Database string `ini:"database"`
	Password string `ini:"password"`
}

type Config struct {
	MysqlConfig `ini:"mysql"`
	RedisConfig `init:"redis"`
}

func loadIni(fileName string, data interface{}) (err error) {
	// 参数校验
	// 传入的data参数必须是指针类型(因为需要修改对其赋值)
	t := reflect.TypeOf(data)
	if t.Kind() != reflect.Ptr {
		err = errors.New("data param should be a ptr")
		return
	}
	// 传进来的data参数必须是结构体类型
	// 判断指针指向的值是否是一个结构体
	// Elem取指针指向的值
	if t.Elem().Kind() != reflect.Struct {
		err = errors.New("data param should be a struct")
		return
	}

	// 读取文件得到字节数据
	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}
	//将字节转换成字符串,并按照换行切割
	lineSlice := strings.Split(string(b), "\r\n")
	fmt.Printf("%#v\n", lineSlice)
	var strcutName string
	// 一行一行的都数据
	for idx, line := range lineSlice {
		if len(line) == 0 {
			//跳过空行
			continue
		}
		//去掉首尾空格
		line = strings.TrimSpace(line)
		// 如果是注释就跳过
		if strings.HasPrefix(line, "//") || strings.HasPrefix(line, "#") {
			continue
		}
		// 如果是[开头表示是节（section）
		if strings.HasPrefix(line, "[") {

			if line[0] != '[' || line[len(line)-1] != ']' {
				err = fmt.Errorf("line:%d syntax error", idx+1)
				return
			}
			//如果[]之间的数据长度为0
			sectionName := strings.TrimSpace(line[1 : len(line)-1])
			if len(sectionName) == 0 {
				err = fmt.Errorf("line:%d syntax error", idx+1)
				return
			}
			// 根据字符串section name去data里面，反射找到结构体
			for i := 0; i < t.Elem().NumField(); i++ {
				field := t.Elem().Field(i)
				if field.Tag.Get("ini") == sectionName {
					strcutName = field.Name
					fmt.Println("find struct ", strcutName)
				}
			}
		} else { //字段
			//1.以等号切割字符串等号左侧为key右侧为value
			if strings.Index(line, "=") == -1 || strings.HasPrefix(line, "=") {
				err = fmt.Errorf("line: %d syntax error", idx+1)
				return
			}
			index := strings.Index(line, "=")
			key := strings.TrimSpace(line[:index])
			value := strings.TrimSpace(line[index+1:])
			//2.根据structName去data里面查找对应字段
			v := reflect.ValueOf(data)
			sValue := v.Elem().FieldByName(strcutName)
			sType := sValue.Type() //拿到嵌套结构体类型
			// structObj := sValue.FieldByName(strcutName)
			if sType.Kind() != reflect.Struct {
				err = fmt.Errorf("%s field not a struct", strcutName)
				return
			}
			//3.遍历嵌套结构体的每一个字段，判断tag是不是等于这个key
			var fieldName string
			var fieldType reflect.StructField
			for i := 0; i < sValue.NumField(); i++ {
				field := sType.Field(i)
				fieldType = field
				if field.Tag.Get("ini") == key { // 找到对应字段
					fmt.Printf("Find %s field\n", fieldName)
					fieldName = field.Name
					break
				}
			}
			//4.如果key=tag，给这个字段赋值
			//4.1根据fieldName去取出对应字段然后赋值
			// fieldObj := sValue.Elem().FieldByName(fieldName)
			// 赋值（根据字段类型赋值）

			fileObj := sValue.FieldByName(fieldName)
			fmt.Printf("AAAAAAAAAAAAAAA:struct %v %s %v\n", strcutName, fieldName, fileObj)
			fmt.Printf("########%v %v\n", fieldName, fieldType.Type.Kind())
			switch fieldType.Type.Kind() {
			case reflect.String:
				fileObj.SetString(value)
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				var valueInt int64
				valueInt, err = strconv.ParseInt(value, 10, 64)
				if err != nil {
					return
				}
				fileObj.SetInt(valueInt)
			}

		}

	}

	return
}

func main() {
	var cfg Config
	err := loadIni("./config.ini", &cfg)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(cfg.MysqlConfig.Address, cfg.MysqlConfig.Port)
}
