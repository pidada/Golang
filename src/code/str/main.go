package main

import (
	"fmt"
	"strings"
)

func main()  {
	s1 := "hello"   //一个英文字母占据一个字节
	s2 := "go语言"  //一个汉字占3个字节

	//求长度
	fmt.Println(len(s1), len(s2))

	//拼接
	fmt.Println(s1 + s2)
	s3 := fmt.Sprintf("%s - %s", s1, s2)  //通过Sprintf方法
	fmt.Println(s3)

	//分割
	s4 := "how do you do"
	fmt.Println(strings.Split(s4, " "))  //两个参数：待分割字符和指定分割的符号
	fmt.Printf("%T\n", strings.Split(s4, " "))  //判断字符串类型

	//包含与否
	fmt.Println(strings.Contains(s4, "do"))

	//判断前后缀
	fmt.Println(strings.HasPrefix(s4, "how"))
	fmt.Println(strings.HasSuffix(s4, "how"))

	//子字符串的位置，通过strings.Index
	fmt.Println(strings.Index(s4, "do"))
	fmt.Println(strings.LastIndex(s4, "do"))

	//join操作
	s5 := []string{"how", "do", "you", "do"}
	fmt.Println(s5)
	fmt.Println(strings.Join(s5, "+"))
}