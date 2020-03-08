### map定义

```
var m1 map[string]int

m1 := make(map[string]int, 10)
```

m1["lisi"] = 100

//特殊写法

```go
v, ok := m1["西二旗"]
if ok {
    fmt.Prinln("m1中存在西二旗")
} else {
    fmt.Println("不存在")
}
```

### 反射

反射是个啥？

优点：

​	让代码更灵活

缺点：

​	影响执行效率。

###### 应用

各种web框架，配置文件解析、ORM框架

###### 在reflect包中的两个方法

TypeOf：

​	可以获取任意对象的类型

​	v := reflect.TypeOf(x)

​	v.Name() v.Kind()

ValueOf





## 接口补充

// 接口由两部分组成： 类型和值

var x interface{} // <Type value>

var a int64 = 12

x  =a // <int64 12>

value,ok := x.(int)	//类型断言

#### 日志包

```
import "github.com/astaxie/beego/logs"


####kafka
package main
import "github.com/Shopify/sarama"

func main(){
    config := sarama.NewConfig()
    config.Producer.RequiredAcks = sarama.WaitForAll
    config.Producer.Partitioner = sarama.NewRandomPartitioner
    config.Producer.Return = true
    
    client, err := sarama.NewSyncProducer([]string{"172.16.6.191:9292"}. config)
}
if err != nil {
    return
}
defer client.Close()
msg := &carama.ProducerMessage{}
msg.Topic = "nginx_log"
msg.Value = sarama.StringEncoder("this is a good test")
pid, offset, err := client.SendMEssage(msg)
if err != nil {
    return
}
fmt.Printf("pid%v offset:%v\n", pid, offset)
}




```



```

```

###tail























