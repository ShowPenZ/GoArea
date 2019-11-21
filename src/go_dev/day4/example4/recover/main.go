/*
 * @Author: ShowPenZ
 * @Date: 2019-11-20 17:32:37
 * @LastEditors: ShowPenZ
 * @LastEditTime: 2019-11-20 17:39:30
 * @Description: recover
 */
package main

import (
	"fmt"
	"time"
)

func test() {
	defer func() {
		err := recover()

		if err != nil {
			fmt.Println(err)
		}
	}()

	b := 0
	a := 100 / b
	fmt.Println(a)
}

func main() {
	for {
		test()
		time.Sleep(time.Second)
	}
}
