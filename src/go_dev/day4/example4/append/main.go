/*
 * @Author: ShowPenZ
 * @Date: 2019-11-20 17:17:31
 * @LastEditors: ShowPenZ
 * @LastEditTime: 2019-11-20 17:19:43
 * @Description: append
 */
package main

import "fmt"

func main() {
	var a []int

	a = append(a, 10, 34, 123)
	a = append(a, a...)
	fmt.Println(a)
}
