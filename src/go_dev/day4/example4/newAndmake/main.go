package main

/*
 * @Author: ShowPenZ
 * @Date: 2019-11-20 16:48:59
 * @LastEditors: ShowPenZ
 * @LastEditTime: 2019-11-21 10:29:00
 * @Description: new用来分配一个内存地址比如int struct 返回的是指针
 */

import "fmt"

func test() {
	var i int
	fmt.Println(i)

	j := new(int)
	// j返回的是指针 指向分配的地址
	fmt.Println(j)

	k := new(int)
	*k = 100
	//*k 指向值
	fmt.Println(*k)

	s1 := new([]int)

	s2 := make([]int, 10)

	*s1 = make([]int, 5)

	(*s1)[0] = 100
	s2[0] = 50

	fmt.Println(s1)
	fmt.Println(s2)

	return
}

func main() {
	test()
}
