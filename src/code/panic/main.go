package main

import "fmt"

func a(){   //定义匿名函数
	fmt.Println("func a")
}

func b(){
	//recover一定要在panic之前，defer和recover联合使用
	defer func(){
		//如果发生了panic错误，通过recover()来恢复
		err := recover()
		if err != nil{   //如果err不是空的，说明err发生错误
			fmt.Println("func b error")
		}
	}()   //()表示对匿名函数的调用
	panic("panic in b")    //panic函数表示错误
}

func c(){
	fmt.Println("func c")
}


func main(){
	a()
	b()
	c()
}