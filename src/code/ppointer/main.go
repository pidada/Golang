package main

import "fmt"

func main(){
	//var a int = 1
	//var ptr1 *int = &a
	//var ptr2 **int = &ptr1
	//var ptr3 **(*int) = &ptr2 // 也可以写作：var ptr3 ***int = &ptr2

	//var a int = 20   //变量定义

	var a int = 20   //定义变量a
	var ptr *int     //定义指针变量ptr
	var pptr **int   //定义指向指针的指针变量

	a = 20

	ptr = &a   //取出a的指针地址
	pptr = &ptr  //取出指针的指针地址

	fmt.Printf("变量 a= %d\n", a)
	fmt.Printf("指针变量 *ptr = %d\n", *ptr)   //*是&的取反操作，得到指针的值
	fmt.Printf("指向指针的指针变量 **pptr = %d\n", **pptr)   //得到指向指针地址的指针的值
}
