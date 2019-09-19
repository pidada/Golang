package main

import (
	"fmt"
)

//二维数组
func main() {
	//cityArray := [...][2]string{   //只有外层可以使用...
	cityArray := [4][2]string{   //总长度是4，里面的长度是2
		{"北京","西安"},
		{"上海","杭州"},
		{"重庆","成都"},
		{"深圳","广州"},
	}
	fmt.Println(cityArray)
	fmt.Println(cityArray[2][0])  //访问某个元素

	//通过for range来遍历所有的元素
	for _, value1 := range cityArray{
		for _, value2 := range value1{
			fmt.Println(value2)
		}
	}
}



