// package main

// import "fmt"

// /* 3.定義 */
// // 演習
// func main() {
// 	/* Q1. 以下の1.11をint型に変換して出力してください。 */
// 	f := 1.11
// 	ff := int(f)
// 	println(ff)

// 	/* Q2. コードを書かずに以下の出力結果を答えてください。 */
// 	s := []int{1, 2, 5, 6, 2, 3, 1}
// 	fmt.Println(s[2:4])

// 	/*
// 		Q3. 以下のコードを実行した時に
// 		fmt.Printf("%T %v", m, m)
// 		以下のような出力結果となるmを作成してください。
// 		map[string]int map[Mike:20 Nancy:24 Messi:30]
// 	*/
// 	m := map[string]int{"Mike": 20, "Nancy": 24, "Messi": 30}
// 	fmt.Printf("%T %v", m, m)

// }

// 21.可変長引数
// func foo(params ...int) {
// 	fmt.Println(len(params), params)
// }

// func main() {
// 	foo(10, 20)
// 	foo(10, 20, 30)
// }

// 20.クロージャー
// func incrementGenerator() func() int {
// 	x := 0
// 	return func() int {
// 		x++
// 		return x
// 	}
// }

// func circleArea(pi float64) func(radis float64) float64 {
// 	return func(radis float64) float64 {
// 		return pi * radis * radis
// 	}
// }

// func main() {
// 	counter := incrementGenerator()
// 	fmt.Println(counter())
// 	fmt.Println(counter())
// 	fmt.Println(counter())
// 	fmt.Println(counter())

// 	c1 := circleArea(3.14)
// 	fmt.Println(c1(2))

// 	c2 := circleArea(3)
// 	fmt.Println(c2(2))
// }

// 19.func
// func add(x, y int) (int, int) {
// 	return x + y, x - y
// }

// func calc(price, item int) (result int) {
// 	result = price * item
// 	return
// }

// func main() {
// 	r1, r2 := add(10, 20)
// 	fmt.Println(r1, r2)

// 	r3 := calc(100, 2)
// 	fmt.Println(r3)

// 	// inner func
// 	f := func(x int) {
// 		fmt.Println("inner func", x)
// 	}
// 	f(1)

// 	func(x int) {
// 		fmt.Println("inner func", x)
// 	}(1)
// }

// 18.byte
// func main() {
// 	b := []byte{72, 73}
// 	fmt.Println(b)
// 	fmt.Println(string(b)) // HI

// 	c := []byte("HI")
// 	fmt.Println(c)
// 	fmt.Println(string(c))
// }

// 17.map
// func main() {
// 	m := map[string]int{"apple": 100, "banana": 200}
// 	fmt.Println(m)
// 	m["new"] = 500
// 	fmt.Println(m)

// 	fmt.Println(m["nothing"])

// 	v, ok := m["apple"]
// 	fmt.Println(v, ok)

// 	v2, ok2 := m["nothing"]
// 	fmt.Println(v2, ok2)
// }

// 15.slice
// func main() {
// 	n := []int{1, 2, 3, 4, 5, 6}
// 	fmt.Println(n)
// 	fmt.Println(n[2])
// 	fmt.Println(n[2:4])
// 	fmt.Println(n[:2])
// 	fmt.Println(n[2:])
// 	fmt.Println(n[:])

// 	n[2] = 100
// 	fmt.Println(n)

// 	n = append(n, 100, 200, 300)
// 	fmt.Println(n)

// 	var board = [][]int{
// 		[]int{0, 1, 2},
// 		[]int{3, 4, 5},
// 		[]int{6, 7, 8},
// 	}
// 	fmt.Println(board)
// }

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
