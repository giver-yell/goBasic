package main

import "fmt"

type Vertex struct {
	X, Y int
}

// 値レシーバ
func (v Vertex) Area() int {
	return v.X * v.Y
}

// ポインタレシーバ
func (v *Vertex) Scale(i int) {
	v.X = v.X * i
	v.Y = v.Y * i
}

func Area(v Vertex) int {
	return v.X * v.Y
}

func main() {
	v := Vertex{3, 4}
	// fmt.Println(Area(v))
	// struct orientedと同様なことがgolagでもできる
	// v.を入力すると、補完でArea()が出てくるので便利
	v.Scale(10)
	fmt.Println(v.Area())
}
