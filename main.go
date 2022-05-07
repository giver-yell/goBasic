package main

import "fmt"

/* 3.定義 */

// 15.slice
func main() {
	n := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(n)
	fmt.Println(n[2])
	fmt.Println(n[2:4])
	fmt.Println(n[:2])
	fmt.Println(n[2:])
	fmt.Println(n[:])

	n[2] = 100
	fmt.Println(n)

	n = append(n, 100, 200, 300)
	fmt.Println(n)

	var board = [][]int{
		[]int{0, 1, 2},
		[]int{3, 4, 5},
		[]int{6, 7, 8},
	}
	fmt.Println(board)
}

// 14.配列
// func main() {
// 	var a [2]int
// 	a[0] = 100
// 	a[1] = 200
// 	fmt.Println(a)

// 	var b [2]int = [2]int{100, 200}
// 	// 配列はサイズを変更できないのでエラー
// 	// b = append(b, 300) // エラー
// 	fmt.Println(b)

// 	// sliceはサイズ変更可能
// 	var c []int = []int{100, 200}
// 	fmt.Println(c)
// }

// 13.型変換
// func main() {
// 	// intからfloat
// 	var x int = 1
// 	xx := float64(x)
// 	fmt.Printf("%T %v %f\n", xx, xx, xx)

// 	// stringからintへの型変換はstrconv.Atoi(A:asciiアスキー)
// 	var s string = "14"
// 	i, _ := strconv.Atoi(s)
// 	fmt.Printf("%T %v\n", i, i)

// 	// ASCII
// 	h := "Hello world"
// 	fmt.Println(string(h[0])) // H
// }

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
