package main //非注释的第一行：声明当前的文件属于哪个包

import "fmt" //导入fmt包，包含格式化输出的函数

func main() {
	//标准声明
	var name string
	var age int
	var isOk bool
	fmt.Println(name, age, isOk)

	//批量声明
	var(
		a string
		b int
		c bool
		d float64
	)
	fmt.Println(a,b,c,d)

	//声明变量指定初始值
	var name1 string = "xiaowangzi"
	var age1 int = 18
	var name3, age3 = "Tom", 28
	fmt.Println(name1, age1, name3, age3)

	//类型推导：系统根据后面的类型自行推导类型；可以在全局中声明
	var name2 = "zhangsan"
	var age2 = 20
	fmt.Println(name2, age2)

	//短变量声明：在函数内部声明
	m := 10
	fmt.Println(m)

	//匿名变量：通过下划线_来实现，不占用命名空间，不会分配内存
}

