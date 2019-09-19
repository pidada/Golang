package main

import "fmt"

func add(x, y int) int{
	return x + y
}

func sub(x, y int) int{
	return x - y
}

//定义3个变量：a,b,op
//a,b：是int类型
//op：是func类型，func接收两个int类型，返回值也是int类型
//calc函数的返回值也是int类型
func calc(x, y int, op func(int, int) int) int{
	return op(x, y)
}


func main(){
	//函数作为参数传递给另一个函数
	r1 := calc(10,20, add)   //30
	fmt.Println(r1)
	r2 := calc(20,10, sub)   //10
	fmt.Println(r2)
}