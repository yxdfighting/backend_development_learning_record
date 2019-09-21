package datastructure_test

import (
	"fmt"
	"testing"
)

type MyInt int64
func TestDataStructure(t *testing.T){
	//测试隐式类型转换,编译报错
	var a int = 6
	//var b float32 = a
	//
	//fmt.Print(a,b)

	//测试强制类型转换,可以编译运行
	var c int64 = int64(a)

	fmt.Print(a,c)

	//测试别名的隐式类型转换
	//var d MyInt = 90
	//
	//var e int64 = d
	//fmt.Print(d,e)
}

func TestPointer(t *testing.T){
	var a int = 8
	var b = &a

	//测测指针运算，无法通过编译
	//c := b + 1

	fmt.Print(a,b,"\n")
	fmt.Printf("%T,%T",a,b)
}

func TestString(t *testing.T){
	var a string

	fmt.Printf("the value of a:%v\n",a)
	fmt.Print(len(a),"\n")

	if a == ""{
		fmt.Print("yes")
	}else{
		fmt.Print("no")
	}
}

func TestCompare(t *testing.T){
	//数组的比较
	a := [...]int{1,2,3,4}
	b := [...]int{1,2,3,5}
	//c := [...]int{1,2,3,4,5}
	d := [...]int{1,2,3,4}

	fmt.Print(a == b,"\n")
	//fmt.Print(a == c)
	fmt.Print(a == d)
}

func TestLoop(t *testing.T){
	for i:=0;i<5;i++{
		fmt.Print(i)
	}
	fmt.Print("\n")

	j := 0
	for j<5{
		j++
		fmt.Print(j)
	}
}

func TestCondition(t *testing.T){
	if a := test();a{
		fmt.Print("测试成功")
	}
}

func test() bool{
	return true
}

func TestSwitch(t *testing.T){
	switch a:= "mmp"; a{
	case "yxd","mmp":
		fmt.Print("yxd")
	case "ydx":
		fmt.Print("ydx")
	default:
		fmt.Print("lly")
	}
}
