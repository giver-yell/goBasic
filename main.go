package main

import (
	"fmt"
)

/* セクション6: Structオリエンテッド */

// 43.インターフェースとダックタイピング
type Human interface {
	Say() string
}

type Person struct {
	Name string
}

type Dog struct {
	Name string
}

func (p *Person) Say() string {
	p.Name = "Mr." + p.Name
	fmt.Println(p.Name)
	return p.Name
}

func DriveCar(human Human) {
	if human.Say() == "Mr.Mike" {
		fmt.Println("run")
	} else {
		fmt.Println("Get out")
	}
}

func main() {
	var mike Human = &Person{"Mike"}
	var x Human = &Person{"x"}
	// mike.Say()
	DriveCar(mike)
	DriveCar(x)
	// DriveCar(Dog)
}

// 42.non-structのメソッド
// あまり使わないのでスキップ

// 41.Embedded
// type Vertex struct {
// 	x, y int
// }

// func (v Vertex) Area() int {
// 	return v.x * v.y
// }

// func (v *Vertex) Scale(i int) {
// 	v.x = v.x * i
// 	v.y = v.y * i
// }

// type Vertex3D struct {
// 	Vertex
// 	z int
// }

// func (v Vertex3D) Area3D() int {
// 	return v.x * v.y * v.z
// }

// func (v *Vertex3D) Scale3D(i int) {
// 	v.x = v.x * i
// 	v.y = v.y * i
// 	v.z = v.z * i
// }

// func New(x, y, z int) *Vertex3D {
// 	return &Vertex3D{Vertex{x, y}, z}
// }

// func main() {
// 	v := New(3, 4, 5)
// 	// v.Scale(10)
// 	fmt.Println(v.Area())
// 	v.Scale3D(10)
// 	fmt.Println(v.Area3D())
// }

// 40.コンストラクタ
// type Vertex struct {
// 	x, y int
// }

// // 値レシーバ
// func (v Vertex) Area() int {
// 	return v.x * v.y
// }

// // ポインタレシーバ
// func (v *Vertex) Scale(i int) {
// 	v.x = v.x * i
// 	v.y = v.y * i
// }

// func New(x, y int) *Vertex {
// 	return &Vertex{x, y}
// }

// func main() {
// 	v := New(3, 4)
// 	v.Scale(10)
// 	fmt.Println(v.Area())
// }

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
