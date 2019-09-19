package main

import "fmt"

//元素类型为：map
func main(){
	//元素类型为map的切片
	var mapSlice = make([]map[string]int, 8, 8)  //完成切片的初始化：map[string]int定义为map类型

	fmt.Println(mapSlice[0] == nil)
	//内部map的初始化
	mapSlice[0] = make(map[string]int, 0)  //map的初始化
	mapSlice[1] = make(map[string]int, 0)  //切片中的每个map都需要进行初始化才能使用
	mapSlice[0]["张三"] = 1000
	mapSlice[1]["李四"] = 100
	fmt.Println(mapSlice)
}
