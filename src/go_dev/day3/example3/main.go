package main

import "fmt"

func modify(p *int) {
	*p = 100000
	return
}

func main() {
	var a int = 5   //定义a的值为5
	fmt.Println(&a) //打印出a的值地址

	var p *int      //定义p为一个指针
	p = &a          //将a的值的地址赋值给p指针
	fmt.Println(*p) //*p 指向p指针的值 *p = 5

	*p = 100       //修改指针所指的值 会一同改变相同值地址的值
	fmt.Println(a) //a = 100

	var b int = 999
	p = &b
	*p = 5

	fmt.Println(a)
	fmt.Println(b)

	modify(&a)
	fmt.Println(a)
}
