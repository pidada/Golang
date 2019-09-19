package main

import (
	"fmt"
)

func main(){
	// 基本写法
	var score = 65
	if score >= 90{   //左花括号和if必须在一行中
		fmt.Println("A")
	} else if score > 75{
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}

	//2.特殊写法
	if score := 65;  score >= 90{   //score变量只在if语句内部生效
		fmt.Println("A")
	} else if score > 75{
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}

}


