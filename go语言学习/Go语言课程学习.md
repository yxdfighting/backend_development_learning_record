## Go语言课程学习

### 第一个go程序

go version 在1.8以后默认设置GoPath

<img src="/home/yxd/.config/Typora/typora-user-images/1568904305025.png" alt="1568904305025" style="zoom:50%;" div align="left"/>

![1568904551617](/home/yxd/.config/Typora/typora-user-images/1568904551617.png)

```go
package main

import (
	"fmt"
	"os"
)

func main(){
	fmt.Println("Hello World")
	fmt.Println(os.Args)
}
```
 可以直接运行go run hello_world.go，也可以编译源码go build hello_world.go,此时会生成一个二进制文件

<img src="/home/yxd/.config/Typora/typora-user-images/1568906026480.png" alt="1568906026480" style="zoom:50%;" div align="left"/>

应用程序的入口：

必须package main,只有main package是可执行的二进制包，

每个文件夹表示一个包，但是文件夹名称可以与package不同

必须是main方法，文件名可以随意

main函数不支持返回值



go语言获取命令行参数通过os.Args获取命令行参数, 可以看到我们打印的部分其实就是helloworld.go编译后的二进制文件 后面的hahaha就是我们的命令行参数

### 变量常量

编写测试程序  

测试文件命名类似,xxx_test.go    

测试方法名命名类似func TestXXX(t *testing.T) {...},大写的方法名在包外可以访问

如下是一个斐波纳契数列的go语言实现

```go
package fib

import (
	"fmt"
	"testing"
)

func TestFib(t *testing.T) {
	//var (
	//	first int =1
	//	second int =1
	//)
	first := 1
	second := 1

	fmt.Print(first,"")
	fmt.Print("",second)
	for m:=0;m<10;m++{
		third := second + first
		fmt.Print("",third)

		first = second
		second = third
	}
	fmt.Print()
}
```

我们通过go test -v fib_test.go可以运行相应的测试文件

<img src="/home/yxd/.config/Typora/typora-user-images/1568910782902.png" alt="1568910782902" style="zoom:50%;" div align="left" />

go是一种强类型的语言，但是go可以根据赋值推断变量类型；

go可以在一个赋值语句中进行多个变量的赋值，比如a,b=b,a

​	

常量定义：

iota在const关键字出现时将被重置为0(const内部的第一行之前)
```go
package const_test

import (
	"testing"
)

const a int = 0

const(
	Monday = iota + 1
	Tuesday
	Wednesday
)

const x = iota // x=0
const (
	y = iota     //y=0
	z            //z=1
)

const (
	m = 1 << iota  //即 1 << 0 把1按照二进制左移0位  00000001
    n                         // 1 << 1                                                     00000010
)

func TestConst(t *testing.T) {
	t.Log(Monday,Tuesday,Wednesday)
}
```

### 数据类型

int int8 int16 int32 int64 uint uint8 uint16 unint32 uint64    int uint 位数与编译器一致

float32 float64

complex64 complex128

bool

string   注意string是值类型，默认的初始值是空字符串，而不是nil

byte 同uint8

rune  同int32



不支持隐式类型转换，不支持别名与原有类型隐式类型转换

很多语言小的类型可以往大的类型进行转换，不会丢失数据，go不支持



典型的预定义值

math.MaxInt64

math.MaxFloat64

math.MaxUint32 等等



支持GC和指针访问内存空间，但是注意不支持指针运算



### 运算符

算术运算符：++a  --b 不允许

比较运算符： 根据具体的类型确定比较规则，比如数组 长度不相同不可以进行比较	

位运算符：按位清零运算符  	

&^   如果右边对应位是1，那么就把左边对应位清零

如果右边对应位是0，左边对应位保持一致



顺便学习以下go中pkg fmt的几个方法区别

Print  仅仅输出到控制台，不接受格式化参数和操作

PrintIn  输出到控制台并换行

Printf    仅可以输出格式化的字符串类型

 %v  按值的原始类型输出     %+v  输出对应key值	





### 循环与条件

go中循环仅有关键字for

if 语句中支持变量赋值，比如
```go
func TestCondition(t *testing.T){
	if a := test();a{
		fmt.Print("测试成功")
	}
}

func test() bool{
	return true
}
```

switch语句,如下:
```go
func TestSwitch(t *testing.T){
	switch a:= "yxd"; a{
	case "yxd":
		fmt.Print("yxd")
	case "ydx":
		fmt.Print("ydx")
	default:
		fmt.Print("lly")
	}
}
```

区别：

- switch的条件可以是各种类型
- 每个case可以不加break，此时仍然不会像其他语言一样每个case都执行
- 单个case中，如果出现多个选项，使用逗号分割