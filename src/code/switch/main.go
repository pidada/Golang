package main

import "fmt"

func main(){
	//猜手指
	//finger := 3
	//switch finger {
	//case 1:
	//	fmt.Println("大拇指")
	//case 2:
	//	fmt.Println("食指")
	//case 3:
	//	fmt.Println("中指")
	//case 4:
	//	fmt.Println("无名指")
	//case 5:
	//	fmt.Println("小指")
	//default:
	//	fmt.Println("无效输入")

	////分支后面跟多个值
	//number := 5
	//switch  number {
	//case 1,3,5,7,9:
	//	fmt.Println("奇数")
	//case 2,4,6,8,10:
	//	fmt.Println("偶数")
	//default:
	//	fmt.Println("无效输入")
	//}

	//分支后面跟表达式
	number := 5
	switch   {
	case number % 2 == 0 :
		fmt.Println("hello python")
	case number % 2 == 1:
		fmt.Println("hello Go")
	default:
		fmt.Println("无效输入")
	}
}

