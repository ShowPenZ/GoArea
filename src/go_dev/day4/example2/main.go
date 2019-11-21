package main

/*
 * @Author: ShowPenZ
 * @Date: 2019-11-20 10:47:04
 * @LastEditors: ShowPenZ
 * @LastEditTime: 2019-11-20 10:55:34
 * @Description: 回文字符串处理
 */

import "fmt"

func process(str string) bool {
	//rune 等同于int32,常用来处理unicode或utf-8字符
	t := []rune(str)
	for i := range t {
		if i == len(t)/2 {
			break
		}
		last := len(t) - i - 1
		if t[i] != t[last] {
			return false
		}
	}

	return true
}

func main() {
	var str string
	fmt.Scanf("%sd", &str)

	if process(str) {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
}
