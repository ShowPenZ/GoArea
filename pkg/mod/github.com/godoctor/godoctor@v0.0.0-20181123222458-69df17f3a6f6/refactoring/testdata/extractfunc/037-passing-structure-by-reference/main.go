//<<<<<extract,14,2,15,23,foo,pass
package main

import "fmt"

type Pt struct {
	x, y int
}

func main() {
	p := &Pt{3, 4}
	fmt.Println("Old Pt", p)
	p.x = 5
	p.y = 6
	fmt.Print("New Pt", p)
}
