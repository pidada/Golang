package main

import "fmt"

func main(){
	for i := 1; i < 10; i++{
		for j := 1; j < i+1; j++{
			fmt.Printf("%v*%v=%v\t",j,i,i*j)  //j在前面，每次j都是从1开始
		}
		fmt.Println()  // 每次执行一个循环进行一次换行
	}
}
