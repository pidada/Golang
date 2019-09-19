package main

import "fmt"

func main()  {
	s := "a"
	switch {
	case s == "a":
		fmt.Println("a")
		fallthrough   //fallthrough执行满足条件case的下一个case，下条语句仍可以执行
	case s == "b":
		fmt.Println("b")
	case s == "c":
		fmt.Println("c")
	default:
		fmt.Println("...")
	}
}
