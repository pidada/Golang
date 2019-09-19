package main

import "fmt"

//全局变量
var number = 20

//定义函数
func testGlobal() {
	number := 10
	// 可以在函数中使用变量
	// 先在自己函数中进行查找，找到了自己的函数中的变量
	// 函数中没有找到就往外层找，即找全局变量
	fmt.Println("变量number", number)   //变量number 10

	//变量i只在for语句块中生效
	for i := 0; i < 10; i++{
		fmt.Println(i)
	}
	//fmt.Println(i)    外层不能访问内部for语句中的变量
}


func main(){
	//函数能够作为变量
	a := testGlobal
	fmt.Printf("%T\n", a)    //a是func类型
	a()  //直接调用a
}