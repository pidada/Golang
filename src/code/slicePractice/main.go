package main

import (
	"fmt"
	"sort"
)

func main() {
	var a = make([]string, 5, 10)   //定义初始化[     ]  5个空格的空切片
	fmt.Println(a)
	for i := 0; i < 10; i++ {
		a = append(a, fmt.Sprintf("%v", i))
	}
	fmt.Println(a)
	fmt.Printf("%T\n", a)  //[]string类型的切片
	fmt.Println(len(a))  //15


	//请使用内置的sort包对数组var b = [...]int{3, 7, 8, 9, 1}进行排序
	var b = [...]int{4,8,2,9,1}   //定义数组
	sort.Ints(b[:])   //将数组变成了切片，直接排序
	fmt.Println("sort_b:",b)
}

