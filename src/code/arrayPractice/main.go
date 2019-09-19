package main

import "fmt"

func main(){
	var numArray = [5]int{1,3,5,7,8}
	var sum int
	fmt.Println(numArray)

	//for循环索引遍历
	//for i := 0; i < len(numArray); i++ {
	//	sum += numArray[i]
	//}
	//fmt.Println(sum)

	//for range遍历
	for _, value := range numArray{
		sum += value
	}
	fmt.Println(sum)
}