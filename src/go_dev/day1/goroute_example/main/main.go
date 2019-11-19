package main

import (
	"fmt"
	"go_dev/day1/goroute_example/goroute"
)

func main() {
	// for i := 0; i < 100; i++ {
	// 	go goroute.Tgoroute(i)
	// }
	var pipes chan int
	pipes = make(chan int, 1)
	go goroute.Tpipe(100, 300, pipes)

	sum := <-pipes
	fmt.Println(sum)

	// time.Sleep(2 * time.Second)

}
