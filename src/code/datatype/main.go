package main

import (
	"fmt"   // 格式化输出的包
	"math"  //导入math包
)


func main(){
	//十进制--->二进制
	n := 10
	fmt.Printf("%b\n", n)  //二进制b
	fmt.Printf("%d\n", n)  //十进制d

	//八进制，用0开头
	m := 075
	fmt.Printf("%d\n", m)
	fmt.Printf("%o\n", m)

	//十六进制
	f := 0xff
	//fmt.Println(f)
	fmt.Printf("%d\n", f)
	fmt.Printf("%x\n", f)

	// unit8
	var age int8   //0-255，当age=256就会出错
	fmt.Println(age)

	//浮点数
	fmt.Printf("%f\n", math.Pi)
	fmt.Printf("%f\n", math.Phi)
	fmt.Printf("%.2f\n", math.Pi)

	// 查看系统默认最大的浮点数
	fmt.Println(math.MaxFloat32)
	fmt.Println(math.MaxFloat64)

	//复数
	var c1 complex64
	c1 = 1 + 5i
	var c2 complex64
	c2 = 10 - 10i
	fmt.Println(c1, c2)

	//布尔值
	var boo bool
	fmt.Println(boo)  //默认是F
	boo = true
	fmt.Println(boo)

	//字符串
	s1 := "hello beijing\n"
	s2 := "你好，北京"
	fmt.Println(s1, s2)

	//转义字符：打印路径
	//第一个反斜杠是转义字符，第二个表示原样输出反斜杠
	fmt.Println("c:\\code\\go.exe")

	//多行字符串输出：里面的内容是原样输出，不进行任何转义
	s3 := `go语言是谷歌开发的\n
go语言有多种数据类型
多行字符串通过反引号来实现
`
	fmt.Println(s3)
}

