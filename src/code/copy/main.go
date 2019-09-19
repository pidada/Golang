package main

import "fmt"

func main() {
	s1 := make([]int, 3) //[0 0 0]
	s2 := s1             //将s1直接赋值给s2，s1和s2共用一个底层数组
	s2[0] = 100
	fmt.Println(s1) //[100 0 0]
	fmt.Println(s2) //[100 0 0]

	a := []int{1,2,3,4,5}
	b := make([]int, 5, 5)
	c := b     //b直接赋值给c，二者共享数组中的数据，变化同时发生。
	copy(b,a)  //ab之间没有关系，独立
	b[0] = 20
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	//通过copy函数删除元素
	d := []string{"北京","深圳","广州","长沙"}
	d = append(d[0:2], d[3:]...)
	fmt.Println(d)
}