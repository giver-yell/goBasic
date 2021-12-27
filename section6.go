// セクション6 Structオリエンテッド

package main

import "fmt"

// 39. ポインタとポインタレシーバーと値レシーバー
type Vertex struct {
	X, Y int
}

// 値レシーバー（classの代わり）
func (v Vertex) Area() int {
	return v.X * v.Y
}

// ポインタレシーバー
func (v *Vertex) Scale(i int) {
	v.X = v.X * i
	v.Y = v.Y * i
}

// func Area(v Vertex) int {
// 	return v.X * v.Y
// }

func main() {
	v := Vertex{3, 4}
	// fmt.Println(Area(v))
	v.Scale(10)
	fmt.Println(v.Area())
}
