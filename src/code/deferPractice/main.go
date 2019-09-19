package main

import "fmt"

//func f1() int {
//	x := 5
//	defer func (){
//		x++
//	}()
//	return x
//}
//
//func main(){
//	fmt.Println(f1())
//}

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func main() {
	a := 1
	b := 2
	defer calc("1", a, calc("10", a, b))
	a = 0
	defer calc("2", a, calc("20", a, b))
	b = 1
}