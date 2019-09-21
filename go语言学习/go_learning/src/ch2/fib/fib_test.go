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
