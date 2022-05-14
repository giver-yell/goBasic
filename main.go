package main

import "fmt"

/* ステートメント */
// 24.if
func by2(num int) string {
	if num%2 == 0 {
		return "ok"
	} else {
		return "no"
	}
}

func main() {
	result := by2(10)
	if result == "ok" {
		fmt.Println("great")
	}

	if result2 := by2(10); result2 == "ok" {
		fmt.Println("great 2")
	}
}
