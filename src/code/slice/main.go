package main

import "fmt"

//切片
func main(){
	var a []string
	var b = []int{}
	var c = []bool{true, false}
	var d = []bool{false, true}

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(a == nil)    //true
	fmt.Println(b == nil)    //false
	fmt.Println(c == nil)   //false

	//基于数组得到切片
	e := [6]int{1,3,5,6,9,2}
	f := e[2:5]   //切片f是数组e的引用；切片是引用类型
	fmt.Println(f)
	fmt.Printf("%T\n", f)

	//再次切片
	g := e[0:2]
	fmt.Println(g)
	fmt.Printf("%T\n", g)

	//make函数构造切片
	h := make([]int,5,10)  //5代表长度，10代表容量（可省略）
	fmt.Println(h)
	fmt.Printf("%T\n", h)

	//长度和容量
	fmt.Println(len(e))
	fmt.Println(cap(h))

	//赋值拷贝
	i := e
	i[0] = 88
	fmt.Println(e)
	fmt.Println(i)
	fmt.Println(e)
}