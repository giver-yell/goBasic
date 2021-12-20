package main

import "fmt"

// 19.関数
// func add(x int, y int) int {
// 	return x + y
// }

// func main() {
// 	r := add(10, 20)
// 	fmt.Println(r)
// }

func add(x, y int) (int, int) {
	return x + y, x - y
}

func cal(price, item int) (result int) {
	result = price * item
	return // result
}

func main() {
	r1, r2 := add(10, 20)
	fmt.Println(r1, r2)

	r3 := cal(100, 2)
	fmt.Println(r3)

	f := func(x int) {
		fmt.Println("inner func", x)
	}
	f(1)

	func(x int) {
		fmt.Println("inner func", x)
	}(1)
}

// 18. バイト型
// func main() {
// 	b := []byte{72, 73}
// 	fmt.Println(b)
// 	fmt.Println(string(b))

// 	c := []byte("HI")
// 	fmt.Println(c)
// 	fmt.Println(string(c))
// }

// 17.map
// func main() {
// 	m := map[string]int{"apple": 100, "banana": 200}
// 	fmt.Println(m)
// 	fmt.Println(m["apple"])
// 	m["banana"] = 300
// 	fmt.Println(m)
// 	m["new"] = 500
// 	fmt.Println(m)

// 	fmt.Println(m["nothing"])

// 	v, ok := m["apple"]
// 	fmt.Println(v, ok)

// 	v2, ok2 := m["nothing"]
// 	fmt.Println(v2, ok2)

// 	m2 := make(map[string]int)
// 	m2["pc"] = 5000
// 	fmt.Println(m2)

// 	// nil メモリ上にないので
// 	// var m3 map[string]int
// 	// m3["pc"] = 5000
// 	// fmt.Println(m3)

// 	var s []int
// 	if s == nil {
// 		fmt.Println("nil")
// 	}
// }

// 16.スライスのmakeとcap
// func main() {
// n := make([]int, 3, 5)
// fmt.Printf("len=%d cap=%d value=%v\n", len(n), cap(n), n)
// n = append(n, 1, 3)
// fmt.Printf("len=%d cap=%d value=%v\n", len(n), cap(n), n)
// n = append(n, 4, 5, 6)
// fmt.Printf("len=%d cap=%d value=%v\n", len(n), cap(n), n)

// // lenとcapどちらも指定
// a := make([]int, 3)
// fmt.Printf("len=%d cap=%d value=%v\n", len(a), cap(a), a)

// b := make([]int, 0) // 0のスライスをメモリに確保
// var c []int // nill メモリ確保しない
// fmt.Printf("len=%d cap=%d value=%v\n", len(b), cap(b), b)
// fmt.Printf("len=%d cap=%d value=%v\n", len(c), cap(c), c)

// 	c = make([]int, 5)
// 	for i := 0; i < 5; i++ {
// 		c = append(c, i)
// 		fmt.Println(c)
// 	}
// 	fmt.Println(c)

// 	c = make([]int, 0, 5)
// 	for i := 0; i < 5; i++ {
// 		c = append(c, i)
// 		fmt.Println(c)
// 	}
// 	fmt.Println(c)

// }

// 15. スライス
// func main() {
// 	n := []int{1, 2, 3, 4, 5}
// 	fmt.Println(n)
// 	fmt.Println(n[3])
// 	fmt.Println(n[2:4])
// 	fmt.Println(n[:2])
// 	fmt.Println(n[2:])
// 	fmt.Println(n[:])

// 	n[2] = 100
// 	fmt.Println(n)

// 	var board = [][]int{
// 		[]int{0, 1, 2},
// 		[]int{3, 4, 5},
// 		[]int{6, 7, 8},
// 	}
// 	fmt.Println(board)

// 	n = append(n, 100, 200, 300)
// 	fmt.Println(n)
// }

// 14.配列
// func main() {
// 	var a [2]int
// 	a[0] = 100
// 	a[1] = 200
// 	fmt.Println(a)

// 	// 配列はサイズを変更できない
// 	// var b [2]int = [2]int{100, 200}
// 	// // b = append(b, 300) // 配列は値の追加ができない
// 	// fmt.Println(b)

// 	// slice
// 	var b []int = []int{200, 300}
// 	b = append(b, 300)
// 	fmt.Println(b)
// }

// 13.型変換
// func main() {
// 	var x int = 1
// 	xx := float64(x)
// 	fmt.Printf("%T %v %f\n", xx, xx, xx)

// 	var y float64 = 2.1
// 	yy := int(y)
// 	fmt.Printf("%T %v %d\n", yy, yy, yy)

// 	var s string = "14"
// 	// stringをint型に単純には型変換はできない
// 	// i := int(s)
// 	i, _ := strconv.Atoi(s)
// 	fmt.Printf("%T, %v\n", i, i)

// 	h := "Hello world"
// 	fmt.Println(h[0], string(h[0]))
// }

// 12.論理値型
// func main() {
// 	t, f := true, false
// 	fmt.Printf("%T %v %t\n", t, t, t)
// 	fmt.Printf("%T %v %t\n", f, f, 0)

// 	fmt.Println(true && true)
// 	fmt.Println(true && false)
// 	fmt.Println(false && false)

// 	fmt.Println(true || true)
// 	fmt.Println(true || false)
// 	fmt.Println(false || false)

// 	fmt.Println(!true)
// 	fmt.Println(!false)

// }

// 11.文字列型
// func main() {
// 	fmt.Println("Hello world")
// 	fmt.Println("Hello" + "world")
// 	fmt.Println("Hello world"[0])
// 	fmt.Println(string("Hello world"[0]))
// 	fmt.Printf("%T", "Hello"[0])

// 	var s string = "Hello world"
// 	fmt.Println(strings.Replace(s, "H", "X", 1))
// 	fmt.Println(strings.Contains(s, "world"))

// 	fmt.Println("test\n" +
// 		"test")
// 	fmt.Println(`test
// 	test
// 	      test`)

// 	fmt.Println("\"")
// 	fmt.Println(`"`)
// }

// 10.数値型
// func main() {
// var (
// 	u8  uint8     = 255
// 	i8  int8      = 127
// 	f32 float32   = 0.2
// 	c64 complex64 = -5 + 12i
// )

// fmt.Println(u8, i8, f32, c64)

// fmt.Printf("type=%T, value=%v", u8, u8)

// fmt.Println("1 + 1 =", 1+1)
// fmt.Println("10 - 1 =", 10-1)
// fmt.Println("10 / 2 =", 10/2)
// fmt.Println("10 / 3 =", 10/3)
// fmt.Println("10.0 / 3 =", 10.0/3)
// fmt.Println("10 / 3.0 =", 10/3.0)
// fmt.Println("10 % 2 =", 10%2)
// fmt.Println("10 % 3 =", 10%3)

// x := 0
// x++
// fmt.Println(x)
// x--
// fmt.Println(x)

// 	fmt.Println(1 << 0) // 0001
// 	fmt.Println(1 << 1) // 0010
// 	fmt.Println(1 << 2) // 0100
// 	fmt.Println(1 << 3) // 1000

// }

// 定数
// const Pi = 3.14

// const (
// 	Username = "test_user"
// 	Password = "password"
// )

// var big int = 9223372036854775807 + 1 // オーバーフロー
// const Big = 9223372036854775807 + 1

// func main() {
// 	fmt.Println(Pi, Username, Password)
// 	// fmt.Printf("%T", Pi)
// 	fmt.Println(Big - 1)
// }

// func init() {
// 	fmt.Println("init")
// }

// func bazz() {
// fmt.Println("bazz")
// }

// var (
// 	i    int     = 1
// 	f64  float64 = 1.2
// 	s    string  = "test"
// 	t, f bool    = true, false
// )

// func foo() {
// 	xi := 1
// 	xi = 2
// 	xf64 := 1.2
// 	xs := "test"
// 	xt, xf := true, false
// 	fmt.Println(xi, xf64, xs, xt, xf)
// 	fmt.Printf("%T\n", xf64)
// 	fmt.Printf("%T\n", xi)
// }

// func main() {
// 	fmt.Println(i, f64, s, t, f)
// 	foo()

// 	// bazz()
// 	// fmt.Println("hello world", time.Now())
// 	// fmt.Println(user.Current())
// }
