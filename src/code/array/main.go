package main

import "fmt"

//数组相关内容
func main(){
	var a [3]int   //int默认是0值
	fmt.Println(a)

	//数组初始化
	//var testArray = [3]int{0,0,0}  带上等号后面需要显式地将数组的0元素写出来
	var testArray = [3]int{}
	var numArray = [3]int{1, 2}
	var cityArray = [3]string{"深圳", "长沙", "广州"}  //使用指定的初始值来完成初始化工作
	fmt.Println(testArray)
	fmt.Println(numArray)
	fmt.Println(cityArray)

	//不带上数组的长度，编译器自动判断数组长度
	var numArray1 = [...]int{1, 2}
	var cityArray1 = [...]string{"北京", "上海", "深圳"}
	fmt.Println(numArray1)   //[1 2]，直接打印元素
	fmt.Printf("type of numArray1:%T\n", numArray1)   //type of numArray:[2]int；%T：表示打印元素类型
	fmt.Println(cityArray1)  //[北京 上海 深圳]
	fmt.Printf("type of cityArray1:%T\n", cityArray1) //type of cityArray:[3]string

	//通过指定索引的值来初始化数组
	//指定索引为1和3的元素值
	b := [...]int{1: 1, 3: 5}
	fmt.Println(b)
	fmt.Printf("type of b:%T\n", b)

	var langArray = [...]string{1:"golang", 3:"python", 5:"java", 7:"c"}
	fmt.Println(langArray)
	fmt.Printf("type of langArray：%T\n", langArray)

	//数组遍历
	//for 下标进行遍历
	for i :=0; i < len(langArray); i++{
		fmt.Println(langArray[i])
	}

	//for range进行遍历
	for index, value := range langArray{
		fmt.Println(index, value)
	}

	//for _, value := range langArray{   //通过下划线只取到value值，而不要索引
	//	fmt.Println(value)
	//}
}
