package main

import "fmt"

// 50.channel
func goroutine1(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

func goroutine2(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

func main() {
	s := []int{1, 2, 3, 4, 5}
	c := make(chan int) // 15, 15
	go goroutine1(s, c)
	go goroutine2(s, c)
	// 受信が終わるまで待機
	x := <-c
	fmt.Println(x)
	y := <-c
	fmt.Println(y)

}

// 49.goroutine と sync.WaitGroup
// 並列処理
// func goroutine(s string, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	for i := 0; i < 5; i++ {
// 		// time.Sleep(100 * time.Millisecond)
// 		fmt.Println(s)
// 	}
// }

// func nomal(s string) {
// 	for i := 0; i < 5; i++ {
// 		// time.Sleep(100 * time.Millisecond)
// 		fmt.Println(s)
// 	}
// }

// func main() {
// 	// 並列処理
// 	// 先にnomal()の処理が終了すると、gooroutine()が処理されないのでsync.WaitGroupを利用
// 	var wg sync.WaitGroup
// 	wg.Add(1)
// 	go goroutine("hello", &wg)
// 	nomal("world")
// 	wg.Wait()
// }
