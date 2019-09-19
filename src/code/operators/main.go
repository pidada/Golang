package main

import "fmt"

func main(){
	a := 10
	b := 3
	//1、实现加减乘，求商和求余数
	fmt.Println(a+b)
	fmt.Println(a-b)
	fmt.Println(a*b)
	fmt.Println(a/b)   //求商
	fmt.Println(a%b)  //求余数

	a++
	fmt.Println(a)   //通过自加1，变成11
	a--
	fmt.Println(a)   //通过自减1，变成10

	//2、关系运算符
	fmt.Println(10 > 2)
	fmt.Println(10 != 2)
	fmt.Println(4 > 5)

	//3、逻辑运算符
	fmt.Println(10>5 && 1==1)
	fmt.Println(!(10>5))
	fmt.Println(1>5 || 1==1)  //管道符||表示或

	//4、位运算符
	c := 1   //001
	d := 5   //101
	fmt.Println(c & d)  //001--->1
	fmt.Println(c | d)  //101--->5
	fmt.Println(c ^ d)  //100--->4

	fmt.Println(1 << 2)  //1--->100:4
	fmt.Println(4 >> 2)  //100--->1

	fmt.Println(1 << 10) //1024表示容量

	//赋值运算符
	var f int
	f = 5
	f += 5   //f = f + 5
	fmt.Println(f)
	f *= 2
	fmt.Println(f)
}
