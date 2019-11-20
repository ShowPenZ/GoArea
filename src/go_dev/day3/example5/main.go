package main

import "fmt"

type defineFunc func(int, int) int //自定义了一个叫defineFunc类型

func add(a, b int) int {
	return a + b
}

/**
 * @description:
 * @param 1.函数名字后面是它的类型
 * @return:
 */
func operate(op defineFunc, a, b int) int {
	return op(a, b)
}

/**
 * @description:
 * @param {type}
 * @return:
 */
func main() {

	c := add
	sum := operate(c, 100, 100)
	fmt.Println(sum)
}

//map,slice,chan,指针,interface默认以引用方式传递
