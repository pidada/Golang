package main

import (
	"fmt"
)

func main(){
	//声明map类型，没有初始化，a的值就是nil
	var a map[string]int
	fmt.Println(a == nil)  //true
	//map的初始化
	a = make(map[string]int, 8)
	fmt.Println(a == nil)

	//a中添加键值对
	a["深圳"] = 1
	a["北京"] = 2
	//%#v显示出来字符串中的引号
	fmt.Printf("a:%#v\n", a)
	fmt.Printf("type:%T\n", a)

	//声明map的同时并初始化
	b := map[int]bool{
		1:true,
		2:false,
	}
	fmt.Printf("b:%#v\n", b)
	fmt.Printf("type:%T\n", b)

	//判断某个键是否存在
	var scoreMap = make(map[string]int, 8)
	scoreMap["张三"] = 98
	scoreMap["李四"] = 90

	value, ok :=scoreMap["王五"]
	fmt.Println(value, ok)
	if ok{
		fmt.Println("李四在scoreMap中",value)
	} else {
		fmt.Println("查无此人")
	}
}