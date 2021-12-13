package main

import "fmt"

const Pi = 3.14

const (
	Username = "test_user"
	Password = "password"
)

// var big int = 9223372036854775807 + 1 // オーバーフロー
const Big = 9223372036854775807 + 1

func main() {
	fmt.Println(Pi, Username, Password)
	// fmt.Printf("%T", Pi)
	fmt.Println(Big - 1)
}

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
