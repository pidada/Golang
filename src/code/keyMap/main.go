package main

import "fmt"

func main(){
	//值为切片的map：首先定义map，并且初始化
	var sliceMap = make(map[string][]int, 8)  //完成对map的初始化
	v, ok := sliceMap["中国"]   //查看map中是否存在某个键
	if ok{
		fmt.Println(v)
	}else{
		//sliceMap中没有“中国”这个键
		sliceMap["中国"] = make([]int, 8)  //完成对切片的初始化：长度和容量都是8
		sliceMap["中国"][0] = 100
		sliceMap["中国"][2] = 200
		sliceMap["中国"][4] = 400
	}
	//遍历sliceMap
	for k,v := range sliceMap{
		fmt.Println(k,v)
	}
}