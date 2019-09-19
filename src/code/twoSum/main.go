package  main

import "fmt"

func main(){
	fmt.Println("hello world")
	var numArray = [5]int{1, 3, 5, 7, 8}
	for i := 0; i < len(numArray); i++{
		for j := i+1; j <len(numArray);  j++{
			if numArray[i] + numArray[j] == 8{
				fmt.Println(i,j)
			}
		}
	}
}



