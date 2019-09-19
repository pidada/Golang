package main

import (
	"fmt"
	"strings"
)


//makeSuffixFunc函数：参数是suffix，string类型；返回值是func;
//func的参数是string，返回值还是string
//第一个return的返回值是func

func makeSuffixFunc(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {   //判断name的后缀是否是suffix
			return name + suffix  //不是则进行拼接
		}
		return name
	}
}

func main() {
	jpgFunc := makeSuffixFunc(".jpg")   //makeSuffixFunc函数传入string类型的参数
	txtFunc := makeSuffixFunc(".txt")
	fmt.Println(jpgFunc("test")) //test.jpg
	fmt.Println(txtFunc("test")) //test.txt
}
