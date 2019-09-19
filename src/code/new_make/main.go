package main

import "fmt"

func main(){
	var a *int    //只声明了指针变量a并没有初始化
	a = new(int)  //通过new函数进行初始化
	*a = 100
	fmt.Println(*a)


	var b map[string]int   //map类型没有通过make初始化
	b = make(map[string]int, 10)   // 初始化
	b["沙河娜扎"] = 100
	fmt.Println(b)


}
