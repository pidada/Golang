package main

import "fmt"

func main(){
	//匿名函数如何执行
	add := func(x,y int) {   //赋值给变量
		fmt.Println(x + y)
	}
	add(10,20)  //调用执行

	func(x,y int){
		fmt.Println(x + y)
	}(10, 15)  //传入参数立即执行
}