package main

import "fmt"

/* セクション6: Structオリエンテッド */

// 46.カスタムエラー
type UserNotFound struct {
	UserName string
}

func (e *UserNotFound) Error() string {
	return fmt.Sprintf("User not found: %v", e.UserName)
}

func myFunc() error {
	// something wrong
	ok := false
	if ok {
		return nil
	}
	return &UserNotFound{UserName: "Mike"}
}

func main() {
	if err := myFunc(); err != nil {
		fmt.Println(err)
	}
}

// 45.Stringer
// type Person struct {
// 	Name string
// 	Age  int
// }

// func (p Person) String() string {
// 	return fmt.Sprintf("My name is %v.", p.Name)
// }

// func main() {
// 	mike := Person{"Mike", 20}
// 	fmt.Println(mike)
// }

// 44.タイプアサーションとswich type文
// どの型でもOK
// func do(i interface{}) {
// 	/*
// 		ii := i.(int)
// 		ii *= 2
// 		fmt.Println(ii)
// 	*/
// 	switch v := i.(type) {
// 	case int:
// 		fmt.Println(v * 2)
// 	case string:
// 		fmt.Println(v + "!")
// 	default:
// 		fmt.Printf("I don't know %T\n", v)
// 	}
// }

// func main() {
// 	// var i interface{} = 10
// 	// do(i)
// 	do(10)
// 	do("Mike")
// 	do(true)
// }

// 43.インターフェースとダックタイピング
// type Human interface {
// 	Say() string
// }

// type Person struct {
// 	Name string
// }

// type Dog struct {
// 	Name string
// }

// func (p *Person) Say() string {
// 	p.Name = "Mr." + p.Name
// 	fmt.Println(p.Name)
// 	return p.Name
// }

// func DriveCar(human Human) {
// 	if human.Say() == "Mr.Mike" {
// 		fmt.Println("run")
// 	} else {
// 		fmt.Println("Get out")
// 	}
// }

// func main() {
// 	var mike Human = &Person{"Mike"}
// 	var x Human = &Person{"x"}
// 	// mike.Say()
// 	DriveCar(mike)
// 	DriveCar(x)
// 	// DriveCar(Dog)
// }

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
