package main

// セクション5 ポインタ

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
