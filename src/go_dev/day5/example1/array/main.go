/*
 * @Author: ShowPenZ
 * @Date: 2019-11-21 10:56:00
 * @LastEditors: ShowPenZ
 * @LastEditTime: 2019-11-21 14:23:24
 * @Description: file content
 */
package main

import "fmt"

func uselessModify(arr [5]int) {
	arr[0] = 100

	return
}

//有效的修改数组 传指针进去
func usefulModify(arr *[5]int) {
	arr[0] = 100
	// 正式点的写法
	//（*arr)[0] = 100

	return
}

//如何定义数组
func defineArray() {
	var a [5]int = [5]int{1, 2, 3, 4, 5}
	var a1 = [5]int{1, 2, 3, 4, 5}
	var a2 = [...]int{38, 53, 234, 5, 34}
	var a3 = [...]int{1: 100, 2: 34}
	var a4 = [...]string{1: "hello", 2: "world"}

	//定义二维数组
	var b [2][5]int = [...][5]int{{1, 2, 3, 4, 5}, {6, 7, 8, 9, 10}}

	fmt.Println(a)
	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(a3)
	fmt.Println(a4)
	fmt.Println(b)
}

//二维数组的循环
func circleArray() {
	var b [2][5]int = [...][5]int{{1, 2, 3, 4, 5}, {6, 7, 8, 9, 10}}

	for row, v := range b {
		for col, v1 := range v {
			fmt.Printf("{%d,%d}=%d ", row, col, v1)
		}
		fmt.Println()
	}
}

func main() {
	var arr [5]int
	var drr [5]int

	uselessModify(arr)

	//取地址
	usefulModify(&drr)

	fmt.Println(arr)
	fmt.Println(drr)

	defineArray()
	circleArray()
}
