package main

import "fmt"

func main() {
	users := map[string]int{"A": 10, "B": 23, "V": 8}
	key := make([]string, len(users))   //设定key 切片长度 len
	value := make([]int, len(users), 4) //设定value的长度 len
	//value := []int{}

	i := 0
	for kk, vv := range users { //key value分开
		key[i] = kk
		i++
		value = append(value, vv)
	}
	fmt.Println(key, value)
	//空用来变量接收不要的值
	for _, vv := range users {
		fmt.Println(vv)
	}
	for kk, _ := range users {
		fmt.Println(kk)
	}
	for kk := range users {
		fmt.Println(kk)
	}
}

/*
问题：
	什么我定义value的时候用make([]int, len(users))的时候输出是
	[A B V] [0 0 0 10 23 8]
	为啥多了3个0？用[]int{}赋值空的时候才是正常的输出
	[A B V] [10 23 8]
*/
