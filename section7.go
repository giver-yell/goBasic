package main

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
