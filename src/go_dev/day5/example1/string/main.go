/*
 * @Author: ShowPenZ
 * @Date: 2019-11-21 19:59:09
 * @LastEditors: ShowPenZ
 * @LastEditTime: 2019-11-21 20:05:23
 * @Description: file content
 */
package main

import "fmt"

func testString() {
	var str string = "hello world"

	s1 := str[0:5]
	s2 := str[6:]

	fmt.Println(s1)
	fmt.Println(s2)
}

func modifyString() {
	var str string = "鹏，hello world"

	s := []rune(str) // 修改字符用rune

	s[0] = 'p'
	s1 := string(s)

	fmt.Println(s1)
}

func main() {
	testString()
	modifyString()
}
