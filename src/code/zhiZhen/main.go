package main

import "fmt"

func main(){
	var a int = 10   //声明实际变量
	var ip *int   //声明指针变量
	ip = &a  //取出a的指针

	//取指针地址，%x表示转义
	fmt.Printf("变量地址：%x\n", &a)
	fmt.Printf("变量地址：%x\n", ip)

	//访问指针的值
	fmt.Printf("*ip变量的值：%d\n", *ip)   //*ip根据指针来取值


	//空指针
	//当一个指针被定义后没有分配到任何变量时，它的值为 nil；nil 指针也称为空指针。
	// nil在概念上和其它语言的null、None、nil、NULL一样，都指代零值或空值。
	// 一个指针变量通常缩写为 ptr

	var ptr *int

	fmt.Printf("ptr的值为：%x\n", ptr)   //空指针值为0
	fmt.Printf("ptr的值为：%x\n", &ptr)  //空指针在内存中也是存在内存地址
}