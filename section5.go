package main

import "fmt"

// セクション5 ポインタ
// 36.struct
type Vertex struct {
	// 小文字だとスコープがprivateになる
	// X int
	// Y int
	X, Y int
	S    string
}

func changeVertex(v Vertex) {
	v.X = 1000
}

func changeVertex2(v *Vertex) {
	v.X = 1000
	// structでは同じ意味
	// (*v).X = 1000
}

func main() {
	v := Vertex{1, 2, "test"}
	changeVertex(v)
	fmt.Println(v)

	v2 := &Vertex{1, 2, "test"}
	changeVertex2(v2)
	fmt.Println(v2)
	fmt.Println(*v2)

	// v := Vertex{X: 1, Y: 2}
	// fmt.Println(v)
	// fmt.Println(v.X, v.Y)
	// v.X = 100
	// fmt.Println(v.X, v.Y)

	// v2 := Vertex{X: 1}
	// fmt.Println(v2)

	// v3 := Vertex{1, 2, "test"}
	// fmt.Println(v3)

	// v4 := Vertex{}
	// fmt.Println(v4)

	// var v5 Vertex
	// fmt.Println(v5)

	// // ポインタが返る
	// v6 := new(Vertex)
	// fmt.Println(v6)

	// v7 := &Vertex{}
	// fmt.Println(v7)
}

// 35.newとmakeの違い
// func main() {

// 	s := make([]int, 0)
// 	fmt.Printf("%T\n", s)

// 	m := make(map[string]int)
// 	fmt.Printf("%T\n", m)

// 	ch := make(chan int)
// 	fmt.Printf("%T\n", ch)

// 	var p *int = new(int)
// 	fmt.Printf("%T\n", p)

// 	var st = new(struct{})
// 	fmt.Printf("%T\n", st)

// 	// メモリ確保
// 	// var p *int = new(int)
// 	// fmt.Println(p)
// 	// fmt.Println(*p)
// 	// *p++
// 	// fmt.Println(*p)

// 	// メモリ確保しない
// 	// var p2 *int
// 	// fmt.Println(p2)
// 	// // メモリ確保しないのでエラー
// 	// *p2++
// 	// fmt.Println(p2)
// }

//34.ポインタ
// func one(x int) {
// 	x = 1
// }

// func one2(x *int) {
// 	*x = 1
// }

// func main() {
// 	var n int = 100
// 	one(n)
// 	// 値は変わらない
// 	fmt.Println(n)

// 	one2(&n)
// 	fmt.Println(n)

// 	// fmt.Println(&n)

// 	// var p *int = &n
// 	// fmt.Println(p)

// 	// fmt.Println(*p)
// }
