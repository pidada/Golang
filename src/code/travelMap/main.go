package main

import (
	"fmt"
	"math/rand"
	"sort"
)


func main(){
	var scoreMap = make(map[string]int, 8)
	scoreMap["张三"] = 98
	scoreMap["李四"] = 90
	scoreMap["王五"] = 97
	scoreMap["小明"] = 89

	//同时遍历键值
	for k, v := range scoreMap{
		fmt.Println(k,v)
	}

	//遍历键k
	for key := range scoreMap{
		fmt.Println(key)
	}

	//遍历value；其中_表示匿名变量
	for _, value := range scoreMap{
		fmt.Println(value)
	}

	//删除指定的键值对
	delete(scoreMap, "小明")
	fmt.Println(scoreMap)


	//固定顺序进行遍历map
	var scoreMap1 = make(map[string]int, 100)
	//添加50个元素
	for i :=0; i < 50; i++{
		key := fmt.Sprintf("stu%02d", i)
		value := rand.Intn(100)
		scoreMap1[key] = value
	}
	for k,v := range scoreMap1{
		fmt.Println(k,v)
	}

	//按照key从小到大的顺序去遍历scoreMap1
	//整体思路：先把key取出来放进切片中；再把该切片进行排序；最后遍历排序后的切片

	//1.取出所有的key放入切片中
	keys := make([]string, 0, 100)   //长度是0，等待放入元素
	for k := range scoreMap1{
		keys = append(keys, k)
	}
	//2.对keys进行排序
	sort.Strings(keys)  //keys目前是个有序的切片
	//3.按照排序后的key对scoreMap排序
	for _, key := range keys{
		fmt.Println(key, scoreMap1[key])
	}
}