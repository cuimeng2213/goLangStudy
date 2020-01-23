## Day08

## 内容回顾：

### 并发之goroutine：

`在关键字go后面写入需要执行的函数即可开启goroutine，函数执行结束goroutine结束`

### 修改日志模块异步写日志功能:



## 今日内容：

### sync包:

`互斥锁`

```
//互斥锁
//sync.Mutex
var lock sync.Mutex
var x int
func f1(){
	for{
		//给公用变量加锁
		lock.Lock()
		x+=1
		lock.Unlock()
	}
}
```

`读写锁：`

当读操作远多于写操作时，使用读写锁效率更更高。

`sync.Once`:

仅执行一次

```
var loadIconOnce sync.Once
func loadIcons(){
icons = map[string]image.Image{
	"left":loadIcon("left.png"),
	}
}
func Icon(name string) image.Image{
	//执行Once
	loadIconOne.Do(loadIcons)
}

```

`sync.Map`:

内置的map不是线程安全的，如果在多线程编程访问公用的map变量可使用sync.Map。

开箱即用，不用make初始化。

`原子操作`：

```
import(
	"sync/atomic"
)
var wg sync.WaitGroup

func add(){
	x++ // 不加锁，或者不使用原子操作，得出的最终值不是预期的值。
}
func main(){
	for i:=0; i<100000; i++{
		wg.Add(1)
		go add()  
	}
	wg.Wait()
}
```



### context

### 网络编程

主要是基于socoket编程。

#### TCP粘包



