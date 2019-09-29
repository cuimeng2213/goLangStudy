package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"reflect"
	"strconv"
	"strings"
)

type Conf struct {
	// fileName string `conf:"file_name"`
	// filePath string `conf:"file_path"`
	// maxSize  int64  `conf:"max_size"`
	// 教训--- 使用结构体反射时字段名字开头需要大写，不然反射的时候找不到对应字段
	FileName string `conf:"file_name"`
	FilePath string `conf:"file_path"`
	MaxSize  int64  `conf:"max_size"`
}

func fillConfig(c interface{}, fieldName, value string) {
	value = strings.Trim(value, "\r\n") //去掉回车换行符
	confType := reflect.TypeOf(c)
	confValue := reflect.ValueOf(c)

	fmt.Printf("%T %s numField= %d\n", confType.Elem().Kind(), confType.Name(), confValue.Elem().NumField())
	if confType.Elem().Kind() == reflect.Struct {
		fmt.Printf("is struct \n")
	}
	for i := 0; i < confValue.Elem().NumField(); i++ {
		field := confType.Elem().Field(i)
		fmt.Printf(">>>field = %v %v %T \n", field.Name, field.Type.Kind(), field.Tag.Get("conf"))
		if field.Tag.Get("conf") == fieldName {
			if field.Type.Kind() == reflect.String {
				fmt.Printf(">>>>>>:%#v  value=%s \n", confValue.Elem().FieldByName(field.Name), value) // 根据字段名字获取 字段信息
				confValue.Elem().FieldByName(field.Name).SetString(value)
			} else if field.Type.Kind() == reflect.Int64 {
				fmt.Printf("AAAAAAAAAAAAAAAAA: int %s \n", value)
				valueInt, _ := strconv.Atoi(value)
				fmt.Printf("AAAAAAAAAAAAAAAAA chaneg: int %d \n", value)
				confValue.Elem().Field(i).SetInt(int64(valueInt))
			}

		}
	}
}

func parseConfig(fileName string, c *Conf) {
	file, err := os.Open(fileName)
	reader := bufio.NewReader(file)

	if err != nil {
		return
	}

	for {
		data, err := reader.ReadString('\n')
		fmt.Printf("err is %T %v \n", err, err)
		if err == io.EOF {
			dataSlice := strings.Split(data, "=")
			fmt.Printf("key=%s v=%v %T\n", dataSlice[0], dataSlice[1], dataSlice[1])
			fillConfig(c, dataSlice[0], dataSlice[1])
			break
		}
		if err != nil {
			panic("读取配置文件错误")
		}
		if len(data) == 0 || strings.HasPrefix(data, "#") {
			continue
		}
		if strings.Index(data, "=") == -1 {
			panic("配置文件格式不对")
		}
		dataSlice := strings.Split(data, "=")
		fmt.Printf("key=%s v=%v %T\n", dataSlice[0], dataSlice[1], dataSlice[1])
		fillConfig(c, dataSlice[0], dataSlice[1])

	}
}

func main() {
	var c = &Conf{}
	parseConfig("log.conf", c)
	fmt.Printf(">>>>: [%s] [%d]\n", c.FileName, c.MaxSize)
	fmt.Printf(">>>>: [%s] \n", c.FilePath)
}
