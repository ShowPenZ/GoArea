package main

import "fmt"

/*
 * @Author: ShowPenZ
 * @Date: 2019-11-21 14:24:05
 * @LastEditors: ShowPenZ
 * @LastEditTime: 2019-11-21 15:51:50
 * @Description: 切片是一个数组的引用，所以切片是引用类型
 *				 切片是一个可变长度的数组
 */

//自定义slice strct类型
type slice struct {
	ptr *[100]int
	len int
	cap int
}

/**
 * @description: 自定义make
 * @param {type}
 * @return:
 */

func defineMake(s slice, cap int) slice {
	s.ptr = new([100]int)
	s.cap = cap
	s.len = 0

	return s
}

func modify(s slice) {
	s.ptr[1] = 1000
}

func testSlice2() {
	var s slice

	s = defineMake(s, 10)

	s.ptr[0] = 100
	modify(s)

	fmt.Println(s.ptr)
}

func testSlice3() {
	var a = [10]int{1, 2, 3, 4}

	b := a[1:5]

	fmt.Println(b)
	fmt.Printf("%p\n", b)
	fmt.Printf("%p\n", &a[1])

}

func testSlice() {
	var a []int
	var arr [5]int = [...]int{1, 2, 3, 4, 5}

	a = arr[:]
	fmt.Println(a)
	fmt.Println(cap(a))

	a = arr[2:5]
	fmt.Println(a)
	fmt.Println(cap(a))

	a = arr[0:1]
	fmt.Println(a)
	fmt.Println(cap(a))

}

func main() {
	testSlice3()
}
