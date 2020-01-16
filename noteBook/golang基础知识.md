## for循环

1.标准写法：`fori:=0; i<10; i++{}`

2.省略第一个变量：

`var i int=0`

`for;i<10;i++`

3.省略前后两条语句

`for;i<10;{`

`i++`

`}`

4.死循环：`for {}`

## 基本数据类型

#### 整形

​	无符号整型：`uint8 uint32 uint64`

​	带符号类型：`int int64`

​	int:具体是32位还是64位看操作系统

### 二进制：

​	Go语言中无法直接定义一个二进制数

​	`var n1 = 0777`

​	`fmt.Printf("%o\n")//打印八进制`

#### 浮点数：

​	float32和float64

​	float： go语言中默认是64位

#### 布尔值

​	true和false

​	不能和其他的类型转换

#### 字符串

​	常用方法

​	字符串不能被修改

### 字符串、字符、字节都是什么

字符串：双引号包裹的是字符串

字符：单引号包裹的字符，单个字母、打个字符、单个汉字

字节：1byte = 8Bit

go语言中字符串使用utf-8编码，一个汉字占8个字节

### byte和rune

​	都是类型

​	rune实际是int32的别名

## 流程控制

​	跳出for循环：`break`

​	跳过此次for循环：`continue`

#### switch:

```go
switch n{
	case 1:
	case 2:
	default:
}
```

## 运算符

主要是熟悉各个算数运算符

## 数组

​	存放元素的容器

​	必须指定存放元素的类型和长度

​	数组是值类型：赋值时会创建新的副本，不会改变自身的值。

```go
//声明
var a1 [3]int
var a2 [4]int
fmt.Printf("a1=%T a2=%t\n", a1,a2) //[3]int [4]int
//初始化
//方式1
a1 = [3]int{1,2,3}
//方式2：根据初始值自动推算数组长度
a3 := [...]int{1,2,3,4,5}
//方式3:只初始化一部分
a4:= [5]int{1,2}
//方式4:根据索引初始化
a5:= [5]int{0:1,4:2}
```

#### 数组的遍历

```
citys :=[...]string{"beijing","shanghai"}
//索引遍历
for i:=0; i<len(citys); i++{
    fmt.Println(citys[i])
}
//range遍历
for i,v := range citys {
    fmt.Println(i,v)
}
//多维数组
vara11 [3][2]int
//初始化
a11 = [3][2]int{
    [2]int{1,2},
    [2]int{3,4},
    [2]int{5,6},
}
//多维数组遍历
for _,v := range a11 {
    for j,v2 := range v {
        fmt.Println(j,v2)
    }
}
```



#### 数组练习

```
//求数组[1,3,5,7,8]所有元素值
func sum() int {
    var a1 :=[...]int{1,3,5,7,8}
    var sum int
    for i:=0; i<len(a1); i++{
        sum += a1[i]
    }
    return sum
}
```

```
//找出数组中和为指定的两个元素的下标
//如:[1,3,5,7,8],找出和为8的两个元素下标分别为（0，3）（1，2）
//定义两个for循环，外层的从第一个开始遍历
//内层的for循环从外层后面那个开始找
a := [...]int{1,3,5,7,8}
for i:=0; i<len(a1); i++ {
    for j:=i+1; j <len(a1); j++{
        if a[i]+a[j] == 8 {
            fmt.Println("(",i,j,")")
        }
    }
}
```

## 切片

是一个拥有相同类型元素的可变长度的序列。它是基于数组的一层封装。

切片是引用类型。它的内部包含 地址、长度、容量。

##### 切片的定义

```
//var x []T
var s1 []int
var s2 []string
//切片初始化
s1 = []int{1,2,3}
s2=[]string{"沙河","北京"}
fmt.Println(s1,s2)
```

```
//由数组得到切片
a1 := [...]int{1,3,5,7,9,11}
s3 := a1[0:4] // [1,3,5,6] 左包含又不包含
s4 := a1[1:3] // 3 5
fmt.Println(s3)
// s3的长度和容量分别为多少？
//切片的容量：切片第一个元素所对应的索引到最后
fmt.Println(len(s3), cap(s3)) // len=4 cap=6
fmt.Println(len(s4), cap(s4)) // len=2 cap=5
```

切片再切片呢？

```
s5 := s3[3:] // len=1 cap=3
```

#### 切片的本质:

​	切片就是一个框，框住了一块连续的内存。包含底层数组指针，容量为底层数组长度，切片长度为实际包含数据长度。

#### make:

- make也是用来分配内存的，区别于new， make是用来给切片、channel、map来分配内存。它返回具体的类型的引用。

- new：一般是用来给基本数据类型分配内存的，返回变量的指针地址



```
//make关键字创建切片：make([]T,len, cap)
ss1 := make([]int,0,5)
```

#### 切片不能直接比较

切片之间不能比较。我们不能使用==判断两个切片是否相等。切片唯一合法的比较操作是和nil比较。一个nil值得切片没有底层数组， 一个nil值得切片长度和容量都是0。但是我们不能说一个切片的长度和容量都是0的切片一定是nil，例如下实例：

```
var s1 []int // len(s1)=0 cap(s1)=0 s1 == nil = true
s2 := []int{} //len(s2)=0, cap(s2)=0 s2 == nil = false
s3 := make([]int,0)// len(s3)=0, cap(s3)=0. s3 == nil false
```

切片的赋值拷贝：

​	由于都是指向同一个底层数组，如果修改底层数据，两个切片值均会改变。

```go
s1 := []int{1,2,3,4}
s2:=s1
s1[2]=300
fmt.Println(s1, s2) // s1:[1,2,300,4] s2:[1,2,300,4]
```

切片遍历：

```
//1.索引遍历
for i:=0; i<len(s1); i++{
    fmt.Println(s1[i])
}
//for range 循环遍历
for i, v := range s1 {
    fmt.Println(i,v)
}
```

切片的append和copy（扩容）

```
package main
func main(){
    s1 := []string{"beijing","shanghai","shenzhen"}
    s1 = append(s1,"guangzhou")
    //append 一个切片， 必须使用一个遍历接收返回值，一般使用被扩容的切片变量
    ss := []string{"杭州","江苏"}
    s1 = append(s1, ss...) //... 表示切片拆开
}
```

```
//copy
func main(){
    a1 :=[]int{1,3,5}
    a2 := a1 //赋值
    var a3 = make([]int, 3,5)
    copy(a3, a1) //copy :操作会新开辟一个底层数组，所以修改a1的值a3值不会变
    fmt.Println(a1,a2,a3)
    a1[0] = 100
    fmt.Println(a1,a2,a3)// 100 3 5, 100,3,5 1,3,5
    //从切片中删除元素
    //go语言的的切片不支持删除操作，
    a1 = append(a1[:1], a1[2:]...) //删除a1[1]这个元素 100，5
}
```

```
//append实现原理
func main(){
    x1 :=[...]int{1,3,5}
    //1.切片不保存具体的值
    //2.切片对应一个底层数组
    //3.底层数组都占用一块连续的内存
    s1 := x1[:]
    fmt.Println(s1,len(s1),cap(s1)) // [1,3,5] , 3, 3
    
    s1 = append(s1[:1], s1[2:]...)  //修改了底层数组 【1,5,5】
    
    fmt.Println(s1,len(s1),cap(s1)) // [1,5], 2, 3
    fmt.Println(x1)// ? [1,5,5]
}
```

```
//关于append删除切片中得某个元素
func main(){
    a1 := [...]int{1,3,5,7,9,11,13,15,17}
    s1 := a1[:]
    //删除索引为1的那个3
    //相当于将索引为1后面的数据copy放到索引为1的位置
    s1 = append(s1[:1],s1[2:]...)
    fmt.Println(s1) //{1,5,7,9,11,13,15,17}
    fmt.Println(a1) // {1,5,7,9,11,13,15,17,17}
}
```



#### 扩容策略：

- 首先判断，如果新申请的容量（cap）大于2被的旧容量，最终容量（newcap）就是新申请的容量
- 否则判断，如果旧切片长度小于1024，则最终容量（newcap）就是旧容量的2倍（newcap =doublecap）
- 否则判断，如果旧切片长度大于1024，则最终容量（newcap）从旧容量（oldcap）开始循环增加原来的1/4。即（newcap=oldcap，for{newcap+=newcap/4}）直到最终容量（newcap）大于等新申请的容量（cap）
- 如果最终容量（cap）计算值溢出了，则最终容量就是新申请的容量（newcap）需要

需要注意的是：不同的类型扩容策略还是不一样的。

###### 以下代码输出什么？

```
var a = make([]int, 5, 10)
for i:=0; i<10; i++ {
    a = append(a,i)
}
fmt.Println(a) // ? 输出什么
//【0 0 0 0 0 | 0 1 2 3 4 5 6 7 8 9】
//cap 是多少？ 20
```

## 指针

Go语言中不存在指针操作.只需要记住两个符号：

```
// &：取地址
// *：根据地址取值
func main(){
    n := 18
    p := &n
    fmt.Println(p)
    fmt.Printf("%T\n",p)
    //根据地址取值
    fmt.Printf("%d\n",*p)
}
```

以下代码运行会出现panic, 为什么？

```
var a *int //定义了指针，但是没有申请内存。a是一个空指针.
*a = 100   //此处操作空指针出现panic
fmt.Prrintln(*a)

//初始化指针
var b = new(int)
*b = 200
fmt.Println(*b)
```

## Map

map定义：map[T]T ---> map[int]string

```
func main(){
    // map
    var m1 map[string]int
    
    //m1["lixiang"] = 9000 //还没开辟空间m1为nil 此处panic
    m1 = make(map[string]int, 10) //尽量估算好容器空间，避免程序执行过程中扩容
    m1["lixiang"] = 9000
    fmt.Println(m1["a"]) // 如果获取对应的key值不存在，则返回对应类型的0值
    
    //遍历
    for k,v := range m1 {
        fmt.Printf("k=%v v=%v \n", k, v)
    }
    
    //删除:使用delet关键字
    delete(m1,"lixiang")
    delete(m1,"lixiang2") //删除一个不存在的key， 什么也不做
}
```

go语言的map没有排序功能。所以要按照一定顺序获取map值，需要对key进行排序，然后遍历key值遍历map。

```
rand.Seed(time.Now().UnixNano())
var scoreMap = make(map[string]int, 200)
for i:=0; i<100; i++ {
    key := fmt.Sprintf("stu_%02d", i)
    value := rand.Int(100)
    scoreMap[key]=value
    keys := make([]string, 0,200)
    for k := range socreMap{
        keys = append(keys, k)
    }
    sort(keys)
    for _, key := range keys {
        fmt.Println(scoreMap[key])
    }
}
```

元素为map类型的切片：

```
func main(){
    var s1 = make([]map[int]string,0,100) //元素为map类型的切片
    //s1[0][100] = "A" --此处panic  Index outof range 因为此时切片长度为0
    m1 :=make(map[int]string, 1)
    m1[10] = "A"
    s1 = append(s1,m1 )
    fmt.Println(s1)
}
```

值为切片类型的Map

```
m1 := make(map[int]int,1)
m1[100] = []int{1,3,5}
```

## 函数

```
package main
import (
	"fmt"
)
//函数定义
func sum(a,b int) int {
    return a+b
}

func main(){
    c := sum(12,12)
    fmt.Println(c)
}
```

统计字符个数

```
func main(){
    //1.判断字符串中汉字的数量
    //2.难点是判断一个字符是汉字
    s1 := "Hello沙河"
    //1.依次拿到字符串的字符
    r := 0
    for _, v := range s1{
    	//2.判断这个字符是不是汉字
        //r = append(r,v)
        ret := unicode.Is(unicode.Han, v)
        //3.把汉字出现的次数累加
        if ret == true {
            r +=
        }
    }
  	fmt.Println("sum: ", r)
}
```

回文判断

```
//字符串从左往右读和从右往左读是一样的。
//比如：上海自来水来自海上
ss := "上海自来水来自海上"
/*
解题思路：
	把字符串中的字符拿出来放到一个[]rune切片中
*/
r := make([]rune, len(ss))
for i:=0; i<len(r)/2; i++ {
    if r[i] != r[len(r)-1 - i] {
        return fasle
    }
    return true
}

```

#### 函数的定义和defer

函数多种声明方式：

defer：运行机制

​		defer在panic之前执行

![1579144829751](C:\Users\Administrator\go\src\goLangStudy\noteBook\assets\1579144829751.png)

​	多用于资源释放。比如：关闭文件句柄，socket、

```
// defer
func deferDemo(){
    fmt.Println("start")
    //defer把它后面的语句延迟到函数即将返回的时候再执行
    //多个defer，后进先出
    defer fmt.Println("heiheihei")
    
    fmt.Println("end")
}
func main() {
    deferDemo()
}
// start
// end
// heiheihei
```

经典例子：（面试题）

```
package main

import "fmt"
// Go语言中函数的return不是原子操作，在底层分两步操作的
//1.第一步：返回值赋值
//2.第二步：真正的return返回
// 函数中如果存在defer语句，执行的时机是在第一步和第二步之间：
func f1() int{
    x:=5
    defer func(){
        x++ //修改的是x 不是返回值
    }()
    return x
}
func f2()(x int){ // 声明返回值变量
    x =5
    defer func(){
        x++ //此处修改了返回值
    }()
    return
}
func f3() (y int){
    x :=5
    defer func(){
        x++ //没有修改返回值y
    }()
    return x
}
func f4()(x int){
    defer func(x int){
        x++
    }(x)	//此处是值传递，不会修改返回值
    return 5
}
func main() {
    f1() // 5
    f2() // 6
    f3() // 5
    f4() // 5
}
```

```
// defer 另一个例子
func calc(index string, a,b int) int {
    ret := a+b
    fmt.Println(index, a, b, ret)
    return ret
}
func main() {
    a :=1
    b :=2
    // 此时先执行calc("10",a,b)，然后defer入栈
    // calc("1",1,3)
    defer calc("1",a,calc("10",a,b))
    a =0
    //calc("2",0,2)
    defer calc("2",a, calc("20",a,b))
    b=1
}
// 10 1 2 3
// 20 0 2 2
// 2 0 2 2
// 1 1 3 4
```



#### 变量作用域

```
package main
import (
	"fmt"
)
//全局变量
var x =100
func f1(){
/*
	函数中查找变量顺序：
	1.先在函数内部查找
	2.如果找不到就往函数外部查找，一直找到全局
*/
	//x:=10 //局部变量仅在函数内部可见
    fmt.Println(x)
}
func main(){
    f1() // 100
}
```

#### 函数类型和变量

```
// 函数类型
func f1(){
    fmt.Println("hello")
}
func f2() int {
    fmt.Println("hello")
    return 2
}
//函数也可以作为参数的类型, 可以实现回调功能
func f3( x func()int ) {
    ret := x()
    fmt.Println(ret)
}

//函数可以作为返回值
func f5() func(int, int)int {
    return func(x int, y int)int{
        fmt.Printf("x=%v y=%v \n",x,y)
        return x+y
    }
}
func main(){
    a :=f1
    b :=f2
    fmt.Printf("%T\n", a) // func()
    fmt.Printf("%T\n", a) // func() int
    f3(b)
    c:=f5()
    c(23,23)
    //匿名函数
    func(){
       fmt.Printf("I am niming func") 
    }()
}
```

## 闭包

一个函数包含了外部作用域的一个变量

```
// 闭包应用场景：需求将f2传给f1执行???
// f1是别人写的接口
func f1(f func() ){
	fmt.Println("f1 function")
    f()
}
func f2(x,y int) {
	fmt.Println("f2 function")
    fmt.Println(x+y)
}
// 需求将f2传给f1执行???

//定义一个函数对f2进行包装
func f3(f func(int,int), x,y int) func() {
    return func(){
        fmt.Println("AAAAA: ", x, y)
        //在此处调用f2
        f(x,y)
    }
}

func main() {
    f := f3(f2, 100,200) // 把原来需要传参的f2函数包装成一个不需要传参的函数
    f1(f)
}
```

闭包进阶演示：

```
func makeStuffixFunc(suyffux string) func(string) string{
    return func(name string) string{
        if !string.HasStuffix(name, stuffix){
            return name+stuffix
        }
        return name
    }
}
func main(){
    jpgFunc := makeStuffixFunc(".jpg")
    txtFunc := makeStuffixFunc(".txt")
    fmt.Println(jpgFunc("test")) // test.jpg
    fmt.Println(jpgFunc("test2")) // test2.txt
}
```

## 内置函数

```
len				计算容器长度（切片、数组）
close			关闭资源文件（channel，文件句柄）
new				申请内存，主要分配值类型，
make			分配内存。主要用于切片、channel、map，返回对象本身
append			追加元素到 数组、切片
panic和recover  用来做错误处理
```

#### 作业：

![1579168393567](C:\Users\Administrator\go\src\goLangStudy\noteBook\assets\1579168393567.png)

```
var (
	coins = 50
	users = []string{
        "Matthew","Sarah"
	}
	dispatchbution = make(map[string]int, len(users))
)
// 思路
func DispatchCoins(){
    
}
func main(){
    
}
```

