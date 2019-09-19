package main

import "fmt"

//多返回值支持类型简写
func calc(a, b int)(sum, sub int){
	sum = a + b
	sub = a - b
	return
	//return sum, sub
}

func main(){
	x, y := calc(20,10)
	fmt.Println(x)
	fmt.Println(y)
}


