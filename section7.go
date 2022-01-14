package main

// 演習
/*
Q1. 以下のように文字列を連結させて出力したいコードがありますが、

test1!
test1!test2!
test1!test2!test3!
test1!test2!test3!test4!
以下のコードには間違いがあります。上記の出力になるように修正してください。

package main

import "fmt"

func goroutine(s []string, c chan int){
    sum := ""
    for _, v := range s{
        sum += v
        c <- sum
    }
}

func main(){
    words := []string{"test1!", "test2!", "test3!", "test4!"}
    c := make(chan int)
    go goroutine(words, c)
    for w := range c{
        fmt.Println(w)
    }
}
*/

import "fmt"

func goroutine(s []string, c chan string) {
	defer close(c)
	sum := ""
	for _, v := range s {
		sum += v
		c <- sum
	}
}

func main() {
	words := []string{"test1!", "test2!", "test3!", "test4!"}
	c := make(chan string)
	go goroutine(words, c)
	for w := range c {
		fmt.Println(w)
	}
}

// 57.sync.Mutex
// type Counter struct {
// 	v   map[string]int
// 	mux sync.Mutex
// }

// func (c *Counter) Inc(key string) {
// 	c.mux.Lock()
// 	defer c.mux.Unlock()
// 	c.v[key]++
// }

// func (c *Counter) Value(key string) int {
// 	c.mux.Lock()
// 	c.mux.Unlock()
// 	return c.v[key]
// }

// func main() {
// 	// c := make(map[string]int)
// 	c := Counter{v: make(map[string]int)}
// 	go func() {
// 		for i := 0; i < 10; i++ {
// 			c.Inc("key")
// 		}
// 	}()
// 	go func() {
// 		for i := 0; i < 10; i++ {
// 			c.Inc("key")
// 		}
// 	}()
// 	time.Sleep(1 * time.Second)
// 	fmt.Println(c, c.Value("key"))
// }

// 56.Default Selectionと for break
// func main() {
// 	tick := time.Tick(100 * time.Millisecond)
// 	boom := time.After(500 * time.Millisecond)
// 	// select for を抜ける
// OuterLoop:
// 	for {
// 		select {
// 		case <-tick:
// 			fmt.Println("tick.")
// 		case <-boom:
// 			fmt.Println("boom.")
// 			break OuterLoop
// 			// retrun
// 		default:
// 			fmt.Println(".")
// 			time.Sleep(50 * time.Millisecond)
// 		}
// 	}
// 	fmt.Println("#######")
// }

// 55.channelとselect
// func goroutine1(ch chan string) {
// 	for {
// 		ch <- "packet from 1"
// 		time.Sleep(3 * time.Second)
// 	}
// }

// func goroutine2(ch chan string) {
// 	for {
// 		ch <- "packet from 2"
// 		time.Sleep(1 * time.Second)
// 	}
// }

// func main() {
// 	c1 := make(chan string)
// 	c2 := make(chan string)
// 	go goroutine1(c1)
// 	go goroutine2(c2)

// 	for {
// 		select {
// 		case msg1 := <-c1:
// 			fmt.Println(msg1)
// 		case msg2 := <-c2:
// 			fmt.Println(msg2)
// 		}
// 	}
// }

// 54.fan-out fan-in
// func producer(first chan int) {
// 	defer close(first)
// 	for i := 0; i < 10; i++ {
// 		first <- i
// 	}
// }

// // chanがoutかinか明示的にするパターン
// func multi2(first <-chan int, second chan<- int) {
// 	defer close(second)
// 	for i := range first {
// 		second <- i * 2
// 	}
// }

// func multi4(second chan int, third chan int) {
// 	defer close(third)
// 	for i := range second {
// 		third <- i * 4
// 	}
// }

// func main() {
// 	first := make(chan int)
// 	second := make(chan int)
// 	third := make(chan int)

// 	go producer(first)
// 	go multi2(first, second)
// 	go multi4(second, third)
// 	for result := range third {
// 		fmt.Println(result)
// 	}
// }

// 53.producerとconsumer
// func producer(ch chan int, i int) {
// 	// something
// 	ch <- i * 2
// }

// func consumer(ch chan int, wg *sync.WaitGroup) {
// 	for i := range ch {
// 		// エラー時にwg.Done()が呼ばれない対策としてインナーfunc()を利用
// 		func() {
// 			defer wg.Done()
// 			fmt.Println("process", i*1000)
// 		}()
// 	}
// 	fmt.Println("#######")
// }

// func main() {
// 	var wg sync.WaitGroup
// 	ch := make(chan int)

// 	// Producer
// 	for i := 0; i < 10; i++ {
// 		wg.Add(1)
// 		go producer(ch, i)
// 	}

// 	// Consumer
// 	go consumer(ch, &wg)
// 	wg.Wait()
// 	// close がないとfor文が完了しない
// 	close(ch)
// 	time.Sleep(2 * time.Second)
// 	fmt.Println("Done")
// }

// 52.channelのrangeとclose
// func goroutine(s []int, c chan int) {
// 	sum := 0
// 	for _, v := range s {
// 		sum += v
// 		fmt.Println(sum)
// 	}
// 	close(c)
// }

// func main() {
// 	s := []int{1, 2, 3, 4, 5}
// 	c := make(chan int, len(s))
// 	go goroutine(s, c)
// 	for i := range c {
// 		fmt.Println(i)
// 	}
// }

// 51.Buffered Channels
// func main() {
// 	// 第二引数でバッファの数を指定
// 	ch := make(chan int, 2)
// 	ch <- 100
// 	fmt.Println(len(ch))
// 	ch <- 200
// 	fmt.Println(len(ch))
// 	// close で channelを閉じる。closeがないと、forループで3つ目も取りに行こうとしてエラーになる
// 	close(ch)

// 	for c := range ch {
// 		fmt.Println(c)
// 	}
// 	// 値を取り出した後はlen が0に戻る
// 	fmt.Println(len(ch))

// 	// 3つ目の受信はエラー
// 	// ch <- 300
// 	// fmt.Println(len(ch))
// }

// 50.channel
// func goroutine1(s []int, c chan int) {
// 	sum := 0
// 	for _, v := range s {
// 		sum += v
// 	}
// 	c <- sum
// }

// func goroutine2(s []int, c chan int) {
// 	sum := 0
// 	for _, v := range s {
// 		sum += v
// 	}
// 	c <- sum
// }

// func main() {
// 	s := []int{1, 2, 3, 4, 5}
// 	c := make(chan int) // 15, 15
// 	go goroutine1(s, c)
// 	go goroutine2(s, c)
// 	// 受信が終わるまで待機
// 	x := <-c
// 	fmt.Println(x)
// 	y := <-c
// 	fmt.Println(y)
// }

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
