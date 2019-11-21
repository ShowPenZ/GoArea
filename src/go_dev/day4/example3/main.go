package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
 * @Author: ShowPenZ
 * @Date: 2019-11-20 11:47:46
 * @LastEditors: ShowPenZ
 * @LastEditTime: 2019-11-20 12:01:27
 * @Description: 从输入框读取字符
 */

func count(str string) (wordCount, spaceCount, numberCount, otherCount int) {
	t := []rune(str)
	for _, v := range t {
		switch {
		case v >= 'a' && v <= 'z':
			fallthrough
		case v >= 'A' && v <= 'A':
			wordCount++
		case v == ' ':
			spaceCount++
		case v >= '0' && v <= '9':
			numberCount++
		default:
			otherCount++
		}
	}
	return
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	result, _, err := reader.ReadLine()

	if err != nil {
		fmt.Println("read from console err:", err)
		return
	}

	wc, sc, nc, oc := count(string(result))

	fmt.Printf("word count:%d\n space count:%d\n number count:%d\n other count:%d\n", wc, sc, nc, oc)

}
