package main

import (
	"fmt"
	"os/user"
	"time"
)

// func init() {
// 	fmt.Println("init")
// }

// func bazz() {
// fmt.Println("bazz")
// }

func main() {
	// bazz()
	fmt.Println("hello world", time.Now())
	fmt.Println(user.Current())
}
