package slice_test

import (
	"fmt"
	"testing"
)

func TestArray(t *testing.T){
	var a [3]int
	a[0] = 1

	for idx,res := range a{
		fmt.Print(idx,":",res,"\n")
	}

	i := 0
	b := [5]int{1,2,3,4,5}
	for i<4{
		i ++
		fmt.Print(b[i],"")
	}

	c := [...]float32{0.89,0.89,0.98}
	for _,res := range c{
		fmt.Print(res,"\n")
	}

	//数组切片，左闭右开
	fmt.Print(b[1:len(b)])
	//fmt.Print(b[1:-1])   不支持负数
}
///slice是个结构体，包含三个元素，第一个是个指针，指向元素的连续存储空间。第二个元素是int型，表示元素的个数
///第三个元素cap是内部数组的容量
func TestSliceInit(t *testing.T){
	var s0 []int
	t.Log(len(s0),cap(s0))

	s0 = append(s0, 5)
	t.Log(len(s0),cap(s0))

	s1 := []int{1,2,3,4}
	t.Log(len(s1),cap(s1))


	s2 := make([]int,3,5)
	s2[1] = 5
	s2[2] = 9
	t.Log(len(s2),cap(s2))

	for _,res := range s2{
		t.Log(res)
	}
}

///测试slice的伸缩
///根据输出可以看出增长规律如下，每当len与cap长度相等时，cap就扩大二倍
///同时可以看出每次伴随着slice cap的扩容，其就会重新申请新的内存空间，同时把原来的元素拷贝到新的内存空间
///可以看出打印内部的指针地址时，每当扩容时其地址会发生变化

func TestSliceGrowing(t *testing.T){
	s := []int{}
	var s1 []int

	s1 = append(s1, 1)
	s = append(s, 2)

	for i:=0;i<10;i++{
		s1 = append(s1, 1) //所以slice的append一定是要重新赋值给原来
		t.Log(len(s1),cap(s1))
		fmt.Print(&s1[0],"\n")
	}
}

///切片共享存储结构
func TestSliceShare(t *testing.T){
	weak := []string{"Mon","Thue","Wed","Thur","Fri","Sat","Sun"}

	work := weak[0:5]
	weekday := weak[3:6]

	weak[4] = "hello"

	t.Log(work,weekday)
}