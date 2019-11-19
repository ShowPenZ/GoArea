package goroute

// func Tgoroute(a int) {
// 	fmt.Println(a)
// }

// Tpipe is a function
func Tpipe(a int, b int, c chan int) {
	c <- (a * b)
}
