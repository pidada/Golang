package main

import "fmt"

func main(){
	//byte处理ASCII码；rune处理Unicode编码
	var c1 byte = 'c'
	var c2 rune = 'c'
	fmt.Println(c1, c2)
	fmt.Printf("c1:%T  c2:%T\n", c1, c2)


	s := "hello中国"

	for i := 0; i < len(s); i++{   //遍历方式
		fmt.Printf("%c\n", s[i])
	}

	fmt.Println()  //打印空行分割作用

	for _, r := range s{   // range循环实现
		fmt.Printf("%c\n", r)
	}
	fmt.Println()
}
