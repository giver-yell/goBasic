package main

import "fmt"

/* 5.ポインタ */
// 35. newとmakeの違い
func main() {
	// メモリ領域確保はnew
	var p *int = new(int)
	fmt.Println(*p) // 0
	*p++
	fmt.Println(*p) // 1

	// メモリ領域を確保しない
	var p2 *int
	fmt.Println(p2) //
}

// 34. ポインタ
// func one(x *int) {
// 	*x = 1
// }

// func main() {
// 	var n int = 100
// 	one(&n)
// 	fmt.Println(n)
// 	// fmt.Println(&n)

// 	// var p *int = &n
// 	// fmt.Println(p)
// 	// fmt.Println(*p)
// }

//3. 定義
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
