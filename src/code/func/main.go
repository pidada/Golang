package main

import "fmt"

//函数：没有默认参数
//没有参数和返回值
func sayHello(){
	fmt.Println("hello Peter")
}

//带上参数的函数
func sayHello2(name string){
	fmt.Println("hello", name)
}

//定义参数和返回值的函数
//参数类型简写：func intSum (a, b int) int{
func intSum (a int, b int) int{
	ret := a + b   //声明并且赋值
	return ret   //返回ret
}

//效果同上：指定了返回值的名称和类型，则函数体内不需要进行
func intSum2 (a int, b int) (ret int){
	ret = a + b   //直接使用
	return  //不必写返回，系统自动找到返回值
}

//接收可变参数，参数名后面...表示是可变参数
//可变参数在函数体内是切片类型
func intSum3(a ...int) int {
	//fmt.Printf("%T\n", a)  //查看a的类型
	ret := 0
	for _, arg := range a{
		ret += arg
	}
	return ret
}

//同时接收固定和可变参数，可变参数必须在固定参数后面
func intSum4(a int, b...int) int {
	//fmt.Printf("%T\n", a)  //查看a的类型
	ret := a
	for _, arg := range b{
		ret += arg
	}
	return ret
}


func main(){
	//函数调用
	sayHello()

	n := "小明"
	sayHello2(n)

	r := intSum(20,10)
	fmt.Println(r)

	m := intSum2(20,10)
	fmt.Println(m)

	//根据传入的参数进行计算
	r1 := intSum3()
	r2 := intSum3(10)
	r3 := intSum3(10,20)
	r4 := intSum3(10,20,30)
	fmt.Println(r1)
	fmt.Println(r2)
	fmt.Println(r3)
	fmt.Println(r4)

	//固定+可变参数
	r5 := intSum4(0)  //a=0,b=[]
	r6 := intSum4(10,20) //a=10,b=20
	r7 := intSum4(10,20,30)  //a=10,b=[20,30]
	fmt.Println(r5)
	fmt.Println(r6)
	fmt.Println(r7)
}
