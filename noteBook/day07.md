## Day07

## 内容回顾

#### time包

```
//时间格式
2006-01-02 15：04:05
```

时间类型：

```
import(
	"time"
)
now := time.Now()

```

时间操作:

```
time.After
time.Sub
time.Before
```

#### 反射：

接口类型底层分为两部分：动态类型和动态值

反射机制应用：json徐磊话， orm关系映射

```
//reflect 包
reflect.TypeOf(x)
reflect.ValueOf(x)
```

#### ini解析：

```
//config.ini
[mysql]
address=10.10.10.4
port=3306
username=root
password=roor@root
# 这是一个注释文件
[redis]
host=127.0.0.1
port=6379
password=root
database=0
```

解析配置文件：

```

```



## 今日内容

### 并发

### goutine

### channel

### sync