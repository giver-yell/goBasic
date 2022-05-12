// package main

// import "fmt"

// /* 5.ポインタ */
// // 36.struct
// type Vertex struct {
// 	// 大文字はpublic, 小文字はprivate: methodの中だけで利用可能
// 	X, Y int
// 	S    string
// }

// func changeVertex(v Vertex) {
// 	v.X = 1000
// }

// func changeVertex2(v *Vertex) {
// 	v.X = 1000
// 	// 同じ意味
// 	// (*v).X = 1000
// }

// func main() {
// 	v := Vertex{1, 2, "test"}
// 	changeVertex(v)
// 	fmt.Println(v) // {1 2 test}

// 	v2 := &Vertex{1, 2, "test"}
// 	changeVertex2(v2)
// 	fmt.Println(v2)
// 	fmt.Println(*v2)
// 	/*
// 		v := Vertex{X: 1, Y: 2}
// 		fmt.Println(v)
// 		fmt.Println(v.X, v.Y)
// 		v.X = 100
// 		fmt.Println(v.X, v.Y)

// 		v2 := Vertex{X: 1}
// 		fmt.Println(v2)

// 		v3 := Vertex{1, 2, "test"}
// 		fmt.Println(v3)

// 		v4 := Vertex{}
// 		fmt.Println(v4)

// 		var v5 Vertex
// 		fmt.Println(v5) // Vertexはnilにならない

// 		v6 := new(Vertex)
// 		fmt.Println(v6)

// 		// newと同じ意味: ポインタがわかりやすいのでこちらを利用する場合が多い
// 		v7 := &Vertex{}
// 		fmt.Println(v7)
// 	*/
// }

// 35. newとmakeの違い
// func main() {
// 	// make: ポインタを返さない
// 	// slice
// 	s := make([]int, 0)
// 	fmt.Printf("%T\n", s) // []int

// 	// map
// 	m := make(map[string]int)
// 	fmt.Printf("%T\n", m) // map[string]int

// 	// channel
// 	ch := make(chan int)
// 	fmt.Printf("%T\n", ch) // chan int

// 	// new: ポインタを返す
// 	// struct
// 	var st = new(struct{})
// 	fmt.Printf("%T\n", st) // *struct {}

// 	// メモリ領域確保はnew
// 	var p *int = new(int)
// 	fmt.Println(p)  // 0x1400012c020
// 	fmt.Println(*p) // 0
// 	*p++
// 	fmt.Println(*p) // 1

// 	// メモリ領域を確保しない
// 	var p2 *int
// 	fmt.Println(p2) // nil
// }

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
