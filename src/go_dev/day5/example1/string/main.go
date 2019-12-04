/*
 * @Author: ShowPenZ
 * @Date: 2019-11-21 19:59:09
 * @LastEditors: ShowPenZ
 * @LastEditTime: 2019-11-22 18:20:12
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

func test() {
	var a = [...]int{1, 2, 3, 4, 5}

	b := a[1:3]
	b[0] = 100
	c := append(b, 7)
	b[1] = 100
	fmt.Println(a, c)

}

func main() {
	// testString()
	// modifyString()
	test()
}
