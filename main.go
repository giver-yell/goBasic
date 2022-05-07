package main

import (
	"fmt"
	"strconv"
)

/* 3.定義 */
// 13.型変換
func main() {
	// intからfloat
	var x int = 1
	xx := float64(x)
	fmt.Printf("%T %v %f\n", xx, xx, xx)

	// stringからintへの型変換はstrconv.Atoi(A:asciiアスキー)
	var s string = "14"
	i, _ := strconv.Atoi(s)
	fmt.Printf("%T %v\n", i, i)

	// ASCII
	h := "Hello world"
	fmt.Println(string(h[0])) // H
}

// 7.import
// func main() {
// 	fmt.Println("Hello world", time.Now())
// 	fmt.Println(user.Current())
// }

// 6.Hello world
// func init() {
// 	fmt.Println("Init!")
// }

// func bazz() {
// 	fmt.Println("Bazz")
// }

// func main() {
// 	bazz()
// 	fmt.Println("Hello, world")
// }
