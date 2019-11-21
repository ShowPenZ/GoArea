/*
 * @Author: ShowPenZ
 * @Date: 2019-11-21 10:21:15
 * @LastEditors: ShowPenZ
 * @LastEditTime: 2019-11-21 10:38:56
 * @Description: closure 闭包 一个函数和与其相关的引用环境组合而成的实体
 */
package main

import (
	"fmt"
	"strings"
)

func closure() func(int) int {
	var x int

	return func(delta int) int {
		x += delta

		return x
	}
}

func makeSuffix(suffix string) func(string) string {
	return func(name string) string {
		if strings.HasSuffix(name, suffix) == false {
			return name + suffix
		}

		return name
	}
}

func main() {
	var f = closure()

	fmt.Println(f(1))
	fmt.Println(f(100))

	var f1 = makeSuffix(".jpg")
	fmt.Println(f1("gg"))

	var f2 = makeSuffix(".png")
	fmt.Println(f2("mm"))

}
