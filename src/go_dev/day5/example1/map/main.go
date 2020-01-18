package main

import "fmt"

func testMap() {
	//第一种定义的方式
	var a map[string]string = map[string]string{
		"key": "value",
	}

	//第二种定义的方式
	//a := make(map[string]string, 10)
	a["abc"] = "efg"
	a["abc2"] = "efg"
	a["abc1"] = "efg"

	fmt.Println(a)
}

func testMap2() {
	//第一种定义的方式
	a := make(map[string]map[string]string, 100)

	//第二种定义的方式
	//a := make(map[string]string, 10)
	a["key1"] = make(map[string]string)
	a["key1"]["key2"] = "efg"
	a["key1"]["key3"] = "efg"
	a["key1"]["key4"] = "efg"
	a["key1"]["key5"] = "efg"

	fmt.Println(a)
}

func main() {
	//testMap()
	testMap2()
}
