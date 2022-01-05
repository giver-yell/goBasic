// セクション6 Structオリエンテッド

package main

import "fmt"

// 44.タイプアサーションとSwitch type文
func do(i interface{}) {
	// intのみ
	// ii := i.(int)
	// ii *= 2
	// fmt.Println(ii)

	// stringのみ
	// ss := i.(string)
	// fmt.Println(ss + "!")

	// switch type文
	switch v := i.(type) {
	case int:
		fmt.Println(v * 2)
	case string:
		fmt.Println(v + "!")
	default:
		fmt.Printf("I don't know %T\n", v)
	}
}

func main() {
	do(10)
	do("Mike")
	do(true)
}

// 43.インターフェイスとダックタイピング
// type Human interface {
// 	Say() string
// }

// type Person struct {
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
// 	var x Human = &Person{"X"}

// 	// mike.Say()
// 	DriveCar(mike)
// 	DriveCar(x)

// }

// 41.Embedded（埋め込み）
// type Vertex struct {
// 	x, y int
// }

// type Vertex3D struct {
// 	Vertex
// 	z int
// }

// func (v Vertex3D) Area3D() int {
// 	return v.x * v.y * v.z
// }

// func (v *Vertex3D) Scale(i int) {
// 	v.x = v.x * i
// 	v.y = v.y * i
// 	v.z = v.z * i
// }

// func New(x, y, z int) *Vertex3D {
// 	return &Vertex3D{Vertex{x, y}, z}
// }

// func main() {
// 	v := New(3, 4, 5)
// 	v.Scale(10)
// 	fmt.Println(v.Area3D())

// }

// 40. コンストラクタ
/* pythonのclass例
class Vertex(object):
	def __init__(self, x, y):
		self._x = x
		self._y = y

	def area(self):
		return self._x * self._y

	def scale(self, i):
		self._x = self._x * i
		self._y = self._y * i

v = Vertex(3, 4)
v.scale(10)
print(v.area())
*/

// type Vertex struct {
// 	x, y int
// }

// // 値レシーバー（classの代わり）
// func (v Vertex) Area() int {
// 	return v.x * v.y
// }

// // ポインタレシーバー
// func (v *Vertex) Scale(i int) {
// 	v.x = v.x * i
// 	v.y = v.y * i
// }

// func New(x, y int) *Vertex {
// 	return &Vertex{x, y}
// }

// // func Area(v Vertex) int {
// // 	return v.x * v.y
// // }

// func main() {
// 	// v := Vertex{3, 4}
// 	// fmt.Println(Area(v))
// 	v := New(3, 4)
// 	v.Scale(10)
// 	fmt.Println(v.Area())
// }

// 39. ポインタとポインタレシーバーと値レシーバー
// type Vertex struct {
// 	X, Y int
// }

// // 値レシーバー（classの代わり）
// func (v Vertex) Area() int {
// 	return v.X * v.Y
// }

// // ポインタレシーバー
// func (v *Vertex) Scale(i int) {
// 	v.X = v.X * i
// 	v.Y = v.Y * i
// }

// // func Area(v Vertex) int {
// // 	return v.X * v.Y
// // }

// func main() {
// 	v := Vertex{3, 4}
// 	// fmt.Println(Area(v))
// 	v.Scale(10)
// 	fmt.Println(v.Area())
// }
