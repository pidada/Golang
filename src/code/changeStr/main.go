package main

import "fmt"

func main(){
	s1 := "big"
	//强制类型转换
	byteS1 := []byte(s1)
	byteS1[0] = 'p'
	fmt.Println(string(byteS1))

	s2 := "白萝卜"
	byteS2 := []rune(s2)
	byteS2[0] = '红'   //赋值语句
	fmt.Println(string(byteS2))
}