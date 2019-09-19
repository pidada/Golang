package main

import "fmt"

//数组是值类型：传参和赋值不能改变数组
func main(){
	//数组类型和f1函数中规定的相同；:=用于短变量声明
	x := [3][2]int{
		{1,2},
		{3,4},
		{5,6},
	}
	fmt.Println(x)
	f1(x)
	fmt.Println(x)
	y := x
	y[0][0] = 1000
	fmt.Println(x)
}


func f1(a [3][2]int){   //接收长度为3,2的int类型的数组
	a[0][0] = 100
}