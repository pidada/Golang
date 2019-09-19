package main

import "fmt"

//常量
const pi = 3.14
const e = 2.71

//同时设置多个变量
//const (
//	pi = 3.14
//	e = 2.71
//)

//变量的值相同，可以略写；只能在常量中使用
//const (
//	n1 = 10
//	n2
//	n3
//)

const (
	n1 = iota   //0
	n2 = iota   //1
	n3 = iota   //2
	n4 = iota   //3
)


func main(){
	fmt.Println(n1, n2, n3, n4)
}