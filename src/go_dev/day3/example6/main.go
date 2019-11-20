package main

import "fmt"

/**
 * @description: 可变长度的参数列表
 * @param {type}
 * @return:
 */
func add(a int, arg ...int) int {
	var sum int = a

	for i := 0; i < len(arg); i++ {
		sum += arg[i]
	}

	return sum
}

func concact(a string, arg ...string) (result string) {
	result = a

	for i := 0; i < len(arg); i++ {
		result += arg[i]
	}

	return
}

func main() {
	sum := add(10, 23, 234)
	fmt.Println(sum)

	res := concact("hello", " ", "world")
	fmt.Println(res)
}
