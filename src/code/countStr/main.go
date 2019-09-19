package main

import (
	"fmt"
	"strings"
)

//统计字符串中每个单词出现的次数
//"how do  you do"
//func main(){
//	var s = "how do you do"
//	var wordCount = make(map[string]int, 10) //切片的长度和容量都是10
//
//	//1. 字符串中有哪些单词：用字符串的Split方法
//	words := strings.Split(s, " ")
//	//2. 遍历单词做统计
//	for _, word := range words{
//		v,ok := wordCount[word]
//		if ok {
//			//map中有这个单词，次数加1
//			wordCount[word] = v + 1
//		}else {
//			//map中没有这个单词，次数初始化为1
//			wordCount[word] = 1
//		}
//	}
//	for k,v := range wordCount{
//		fmt.Println(k,v)
//	}
//}


func main(){
	var s = "how do you do"
	var wordCount = make(map[string]int, 10)   //定义切片用来存放数据，类似Python中的空字典

	words := strings.Split(s, " ")
	for _, word := range words{
		v,ok := wordCount[word]
		if ok{
			wordCount[word] = v + 1
		} else{
			wordCount[word] = 1
 		}
	}
	for k,v := range wordCount{
		fmt.Println(k,v)
	}
}
