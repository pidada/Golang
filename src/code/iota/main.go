package main

import "fmt"

//每出现一次iota，计数加1
const (
	n1 = iota   //0
	n2 = iota   //1
	n3 = iota   //2
	n4 = iota   //3

	//使用下划线跳过某些值；值是存在的，但不需要用
	_
	n6 = iota   //5
	_

	//iota声明中间插队
	n8 = 28
)

//重新出现再次归0
const n9 = iota


//定义单位数量级
//其中<<表示左移操作符号，1<<10表示1的二进制左移10位，变成10000000000，也就是十进制的1024
//2<<2：表示2的二进制左移2位，10变成1000，十进制的8

const (
	_  = iota  //出现const，iota置为0
	KB = 1 << (10 * iota) //iota=1，1<<10
	MB = 1 << (10 * iota) //iota=2，1<<20
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)

const (
	a, b = iota + 1, iota + 2  //0+1,0+2
	c, d  //等价于c, d = iota + 1, iota + 2，结果为1+1,1+2
	e, f  //2+1, 2+2
)


func main(){
	fmt.Println(n1, n2, n3, n4, n6, n8, n9)
	fmt.Println(a,b,c,d,e,f)
}