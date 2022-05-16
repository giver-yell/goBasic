package main

import "fmt"

/* セクション6: Structオリエンテッド */

// 40.コンストラクタ
type Vertex struct {
	x, y int
}

// 値レシーバ
func (v Vertex) Area() int {
	return v.x * v.y
}

// ポインタレシーバ
func (v *Vertex) Scale(i int) {
	v.x = v.x * i
	v.y = v.y * i
}

func New(x, y int) *Vertex {
	return &Vertex{x, y}
}

func main() {
	v := New(3, 4)
	v.Scale(10)
	fmt.Println(v.Area())
}

// 39.ポインタレシーバと値レシーバ
// type Vertex struct {
// 	X, Y int
// }

// // 値レシーバ
// func (v Vertex) Area() int {
// 	return v.X * v.Y
// }

// // ポインタレシーバ
// func (v *Vertex) Scale(i int) {
// 	v.X = v.X * i
// 	v.Y = v.Y * i
// }

// func Area(v Vertex) int {
// 	return v.X * v.Y
// }

// func main() {
// 	v := Vertex{3, 4}
// 	// fmt.Println(Area(v))
// 	// object orientedと同様なことがgolagでもできる
// 	// v.を入力すると、補完でArea()が出てくるので便利
// 	v.Scale(10)
// 	fmt.Println(v.Area())
// }
