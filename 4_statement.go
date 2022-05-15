// package main

// import "fmt"

/* ステートメント */

// 演習
/*
Q1 . 以下のスライスから一番小さい数を探して出力するコードを書いてください。

l := []int{100, 300, 23, 11, 23, 2, 4, 6, 4}

Q2. 以下の果物の価格の合計を出力するコードを書いてください。

m := map[string]int{
    "apple":  200,
    "banana": 300,
    "grapes": 150,
    "orange": 80,
    "papaya": 500,
    "kiwi":   90,
}
*/
// func main() {
// 	// Q1
// 	l := []int{100, 300, 23, 11, 23, 2, 4, 6, 4}
// 	var min int

// 	for k, num := range l {
// 		if k == 0 {
// 			min = num
// 			continue
// 		}
// 		if min > num {
// 			min = num
// 		}
// 	}
// 	fmt.Println(min)

// 	// Q2
// 	m := map[string]int{
// 		"apple":  200,
// 		"banana": 300,
// 		"grapes": 150,
// 		"orange": 80,
// 		"papaya": 500,
// 		"kiwi":   90,
// 	}

// 	sum := 0
// 	for _, v := range m {
// 		sum += v
// 	}
// 	fmt.Println(sum)
// }

// 31. panicとrecover
// 例外処理。Golangでは、30.エラーハンドリングが推奨されているので基本的には使用しない

// 30.エラーハンドリング
// func main() {
// 	file, err := os.Open("./main.go")
// 	if err != nil {
// 		log.Fatalln("Error")
// 	}
// 	defer file.Close()
// 	data := make([]byte, 100)
// 	count, err := file.Read(data)
// 	if err != nil {
// 		log.Fatalln("Error")
// 	}
// 	fmt.Println(count, string(data))
// }

// 29.log
// func LoggingSettings(logFile string) {
// 	logfile, _ := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
// 	multiLogFile := io.MultiWriter(os.Stdout, logfile)
// 	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
// 	// log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
// 	log.SetOutput(multiLogFile)
// }

// func main() {
// 	LoggingSettings("test.log")

// 	_, err := os.Open("aaa")
// 	if err != nil {
// 		log.Fatalln("Exit", err)
// 	}

// 	log.Println("loggint")
// 	log.Printf("%T %v", "test", "test")

// 	// fatalはexitされる
// 	log.Fatalf("%T %v", "test", "test")
// 	log.Fatalln("error!")

// 	fmt.Println("ok")
// }

// 28.defer
// func foo() {
// 	defer fmt.Println("world foo")

// 	fmt.Println("hello foo")
// }

// func main() {
// 	/*
// 		defer fmt.Println("world")

// 		foo()

// 		fmt.Println("HELLO")
// 	*/
// 	/*
// 		fmt.Println("run")
// 		defer fmt.Println(1)
// 		defer fmt.Println(2)
// 		defer fmt.Println(3)
// 		fmt.Println("succsess")
// 	*/
// 	file, _ := os.Open("./main.go")
// 	defer file.Close()
// 	data := make([]byte, 100)
// 	file.Read(data)
// 	fmt.Println(string(data))
// }

// 27.switch
// func getOsName() string {
// 	return "aj"
// }

// func main() {
// 	// os := getOsName()
// 	switch os := getOsName(); os {
// 	case "mac":
// 		fmt.Println("mac")
// 	case "windows":
// 		fmt.Println("windows")
// 	default:
// 		fmt.Println("default", os)
// 	}

// 	t := time.Now()
// 	fmt.Println(t)
// 	fmt.Println(t.Hour())
// 	switch {
// 	case t.Hour() < 12:
// 		fmt.Println("morning")
// 	case t.Hour() < 17:
// 		fmt.Println("afternoon")
// 	}
// }

// 26.range
// func main() {
// 	l := []string{"python", "go", "java"}

// 	for i := 0; i < len(l); i++ {
// 		fmt.Println(i, l[i])
// 	}

// 	// foreachの代わり
// 	for i, v := range l {
// 		fmt.Println(i, v)
// 	}

// 	for _, v := range l {
// 		fmt.Println(v)
// 	}

// 	// map
// 	m := map[string]int{"apple": 100, "banana": 200}

// 	for k, v := range m {
// 		fmt.Println(k, v)
// 	}
// 	for k := range m {
// 		fmt.Println(k)
// 	}
// 	for _, v := range m {
// 		fmt.Println(v)
// 	}
// }

// 25.for
// func main() {
// 	for i := 0; i < 10; i++ {
// 		if i == 3 {
// 			fmt.Println("continue")
// 			continue
// 		}

// 		if i > 5 {
// 			fmt.Println("break")
// 			break
// 		}
// 		fmt.Println(i)
// 	}

// 	sum := 1
// 	// for ; sum < 10 {
// 	for sum < 10 {
// 		sum += sum
// 		fmt.Println(sum)
// 	}
// 	fmt.Println(sum)

// 	// 無限ループ
// 	for {
// 		fmt.Println("hello")
// 	}
// }

// 24.if
// func by2(num int) string {
// 	if num%2 == 0 {
// 		return "ok"
// 	} else {
// 		return "no"
// 	}
// }

// func main() {
// 	result := by2(10)
// 	if result == "ok" {
// 		fmt.Println("great")
// 	}

// 	if result2 := by2(10); result2 == "ok" {
// 		fmt.Println("great 2")
// 	}
// }
