package main

import (
	"fmt"
	"go_dev/day1/pack_example/calc"
)

func main() {
	sum := calc.Add(100, 200)
	sub := calc.Sub(300, 100)
	fmt.Println(sum, sub)
}
