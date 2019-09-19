package main

import "fmt"

//切片遍历
func main(){
	numArray := []int{1,2,3,4,5,6,7,8}


	for i :=0; i<len(numArray); i++{
		fmt.Println(i, numArray[i])
	}

	fmt.Println()

	for index, value := range numArray{
		fmt.Println(index, value)
	}
}
