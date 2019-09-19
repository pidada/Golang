package main

import "fmt"

func intSum(a, b int)(ret int){
	ret = a + b
	return
}

//defer语句的打印类似于栈的先进后出FILO
func main(){
	fmt.Println("start")
	defer fmt.Println("1")
	fmt.Println("middle")
	defer fmt.Println("2")
	x := intSum(10,20)
	fmt.Println(x)
	defer fmt.Println("3")
	fmt.Println("end")
}


//倒序执行
//func f1() int {
//	x := 5
//	defer func() {
//		x++
//	}()
//	return x
//}
//
//func f2() (x int) {
//	defer func() {
//		x++
//	}()
//	return  5
//}
//
//func f3() (y int) {
//	x := 5
//	defer func() {
//		x++
//	}()
//	return x
//}
//func f4() (x int) {
//	defer func(x int) {
//		x++
//	}(x)
//	return 5
//}
//func main() {
//	fmt.Println(f1())
//	fmt.Println(f2())
//	fmt.Println(f3())
//	fmt.Println(f4())
//}