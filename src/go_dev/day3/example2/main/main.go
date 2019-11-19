package main

import (
	"fmt"
	"time"
)

func testActionTime() {
	time.Sleep(time.Millisecond * 100)
}

func main() {
	start := time.Now().UnixNano()
	testActionTime()
	end := time.Now().UnixNano()

	fmt.Printf("cost:%d us\n", (end-start)/1000)

}
