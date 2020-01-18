/*
 * @Author: ShowPenZ
 * @Date: 2019-11-21 20:08:12
 * @LastEditors: ShowPenZ
 * @LastEditTime: 2019-11-22 18:18:12
 * @Description: sort
 */
package main

import (
	"fmt"
	"sort"
)

func testIntSort() {
	var a = [...]int{1, 5, 4, 6, 8, 3, 2, 532}

	sort.Ints(a[:])

	fmt.Println(a)
}

func testStringSort() {
	var a = [...]string{"abc", "a", "A", "z", "Z", "eeee"}
	sort.Strings(a[:])

	fmt.Println(a)
}

func testFloat() {
	var a = [...]float64{2.3, 0.5, 0.7, 0.2, 13123.3, 0.002}
	sort.Float64s(a[:])

	fmt.Println(a)
}

func main() {
	testIntSort()
	testStringSort()
	testFloat()
}
