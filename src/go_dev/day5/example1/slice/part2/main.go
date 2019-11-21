/*
 * @Author: ShowPenZ
 * @Date: 2019-11-21 16:12:00
 * @LastEditors: ShowPenZ
 * @LastEditTime: 2019-11-21 18:01:17
 * @Description: slice part2
 */
package main

import "fmt"

func testSlice() {
	var a [5]int = [...]int{1, 2, 3, 4, 5}

	s := a[1:]
	fmt.Printf("s=%p a[1]=%p\n", s, &a[1])

	var c = []int{1, 2, 3}
	var b = []int{4, 5, 6}
	//使用内置函数append 操作切片
	c = append(c, b...)
	fmt.Println(c)

	//使用内置函数copy 操作切片
	var d []int = []int{1, 2, 3, 4, 5, 6}
	e := make([]int, 10)
	copy(e, d)

	fmt.Println(e)
}

func main() {
	testSlice()
}
