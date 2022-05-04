package main

import (
	"fmt"
	"os/user"
	"time"
)

/* 3.定義 */

// 7.import
func main() {
	fmt.Println("Hello world", time.Now())
	fmt.Println(user.Current())
}

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
