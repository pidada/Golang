package main

import "fmt"

func main(){
	var a []int     //定义切片未初始化
 	var b = []int{}  // 定义并且初始化

 	printSlice(a)
	//判断a和nil的关系
	if (a== nil){
		fmt.Println("切片a是空的")
	}

	printSlice(b)
	//判断b和nil的关系
	if (b == nil){
		fmt.Println("切片b是空的")
	}
}

//定义printSlice函数，输出长度、容量和切片
func printSlice(x []int){
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}