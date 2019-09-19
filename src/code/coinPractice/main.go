package main

import (
	"fmt"
)


/*
你有50枚金币，需要分配给以下几个人：Matthew,Sarah,Augustus,Heidi,Emilie,Peter,Giana,Adriano,Aaron,Elizabeth。
分配规则如下：
a. 名字中每包含1个'e'或'E'分1枚金币
b. 名字中每包含1个'i'或'I'分2枚金币
c. 名字中每包含1个'o'或'O'分3枚金币
d: 名字中每包含1个'u'或'U'分4枚金币
写一个程序，计算每个用户分到多少金币，以及最后剩余多少金币？程序结构如下，请实现 ‘dispatchCoin’ 函数
*/
var (
	coins = 50
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie",
		"Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
	distribution = make(map[string]int, len(users))   //定义空map，用于存储数据

)

/*
1、先遍历users数组，取出其中的每个user
2、通过for range遍历出user中的每个字母，通过string函数转换成字符类型，再和字母进行比较
3、map类型的distribution初始化为0；比较完毕，对distribution中的每个value值进行增加操作
4、得到distribution：键为单个user，值为增加的数值
5、通过for range 循环遍历distribution中的每个value，所有的value相加赋值给sum
6、得到left = coins - sum
*/


func dispatchCoin() int{
	for _, v := range users {
		// distribution中的每个键的值初始化为0，用于将来按照要求进行相应的增加操作
		distribution[v] = 0
		//对于单个的users遍历
		for i := 0; i < len(v); i++ {
			//将每个user的字母转成字符类型string，并赋值给word
			word := string(v[i])
			if word == "e" || word == "E" {
				distribution[v] += 1
			}
			if word == "i" || word == "I" {
				distribution[v] += 2
			}
			if word == "o" || word == "O" {
				distribution[v] += 3
			}
			if word == "u" || word == "U" {
				distribution[v] += 4
			}
		}
	}
	sum := 0
	//distribution是map类型，存储的是键值对形式，key是单个user，value是上面的number
	for _, v := range distribution {
		sum += v
	}
	left := coins - sum
	return left
}


func main() {
	left := dispatchCoin()
	fmt.Println("剩下：", left)
}
