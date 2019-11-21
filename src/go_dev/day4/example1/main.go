package main

/*
 * @Author: ShowPenZ
 * @Date: 2019-11-20 10:01:44
 * @LastEditors: ShowPenZ
 * @LastEditTime: 2019-11-20 10:12:21
 * @Description: defer
 */

import "fmt"

func main() {
	var i int = 0

	defer fmt.Println(i)
	defer fmt.Println("hello")

	i = 10
	fmt.Println(10)

}
