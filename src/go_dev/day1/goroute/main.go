package main

import "time"

func main() {
	for i := 0; i < 100; i++ {
		go testGoroute(i)
	}

	time.Sleep(2 * time.Second)

}
