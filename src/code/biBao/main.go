package main

import "fmt"

func a() func(){  // 定义a函数，返回值的类型是func
	name := "GO语言"
	return func(){  //返回的是匿名函数
		fmt.Println("hello", name)  //先在匿名函数中找name，没有再到匿名函数外部查找
	}
}

func main(){
	r := a()   //将函数赋值给变量r
	r()   //相当于是执行了a函数里面的匿名函数
}

