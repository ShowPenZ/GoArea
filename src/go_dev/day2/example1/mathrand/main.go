package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for i := 0; i < 10; i++ {
		a := rand.Int()
		b := rand.Intn(100)
		fmt.Println(a)
		fmt.Println(b)
	}

}
