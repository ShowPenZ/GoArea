/*
 * @Author: ShowPenZ
 * @Date: 2019-11-21 10:56:00
 * @LastEditors: ShowPenZ
 * @LastEditTime: 2019-11-21 11:00:45
 * @Description: file content
 */
package main

import "fmt"

func uselessModify(arr [5]int) {
	arr[0] = 100

	return
}

func usefulModify(arr *[5]int) {
	arr[0] = 100
	// 正式点的写法
	//（*arr)[0] = 100

	return
}

func main() {
	var arr [5]int
	var drr [5]int

	uselessModify(arr)

	usefulModify(&drr)

	fmt.Println(arr)
	fmt.Println(drr)
}
