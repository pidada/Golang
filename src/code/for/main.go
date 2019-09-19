package main

import "fmt"

//for循环基本写法
func main(){
	//基本写法
	for i := 0; i < 10; i++{
		fmt.Println(i)
	}
	//for循环省略初始语句，但是必须保留分号
	var j = 0
	for ; j < 10; j++{
		fmt.Println(j)
	}

	//省略初始语句和结束语句
	var k = 10
	for k > 0{
		fmt.Println(k)
		k--
	}


	//无限循环
	//for {
	//	fmt.Println("hello Go!")
	//}

	//break跳出循环
	for i := 0; i < 5; i++{
		fmt.Println(i)
		if i == 3{   //i等于3，跳出循环体，结束for循环
			break
		}
	}
	//continue跳出循环
	for i := 0; i < 5; i++{
		if i == 3{
			continue  //跳过本次for循环，不执行打印语句，直接继续下次循环
		}
		fmt.Println(i)
	}


}

