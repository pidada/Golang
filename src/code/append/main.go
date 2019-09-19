package main

import "fmt"

func main(){
	var a []int      //此时a没有申请内存，不能直接添加元素
	//a[0] = 100
	//fmt.Println(a)

	for i :=0; i<10; i++{
		a = append(a,i)  // append函数将元素追加到切片的最后，同时需要原来的切片进行接收
		fmt.Printf("%v len:%d cap:%d ptr:%p\n", a, len(a), cap(a), a)
	}

	//append追加多个元素
	var citySlice []string
	citySlice = append(citySlice, "北京")   // 追加一个元素
	citySlice = append(citySlice, "上海", "广州", "深圳")  	// 追加多个元素
	b := []string{"成都", "重庆"}
	citySlice = append(citySlice, b...)   	// 追加切片：三个点表示拆包，编译器自动将b中的元素一个个取出来放进a
	fmt.Println(citySlice) //[北京 上海 广州 深圳 成都 重庆]
}
