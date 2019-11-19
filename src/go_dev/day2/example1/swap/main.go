package main

import "fmt"

func swap(a *int, b *int) { //指针指的是值
	tmp := *a
	*a = *b //值交换
	*b = tmp
	return
}

func main() {
	first := 100
	second := 200

	swap(&first, &second) // &first 意思是把first 的地址传进去
	fmt.Println("first=", first)
	fmt.Println("second=", second)
}
