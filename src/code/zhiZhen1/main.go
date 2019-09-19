package main

import "fmt"

//指针传值
func modify1(x int){   //定义变量
	x = 100
}

func modify2(x *int) {  //定义指针变量
	*x = 100
}


func main(){
	a := 10

	modify1(a)
	fmt.Println(a)  //10
	modify2(&a)
	fmt.Println(a)  //100
}