package main

// 30.エラーハンドリング（例外の代わり）
// func main() {
// 	// 返り値が2つ
// 	file, err := os.Open("./section4.go")
// 	if err != nil {
// 		log.Fatalln("Error!")
// 	}
// 	defer file.Close()
// 	data := make([]byte, 100)
// 	count, err := file.Read(data)
// 	if err != nil {
// 		log.Fatalln("Error2!")
// 	}
// 	fmt.Println(count, string(data))

// 	// 返り値が1つなら1行で書く
// 	if err = os.Chdir("test"); err != nil {
// 		log.Fatalln("Error")
// 	}
// }

// // 29.log
// func LoggingSettings(logFile string) {
// 	logfile, _ := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
// 	multiLogFile := io.MultiWriter(os.Stdout, logfile)
// 	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
// 	log.SetOutput(multiLogFile)
// }

// func main() {
// 	LoggingSettings("test.log")
// 	_, err := os.Open("faeifjojdf")
// 	if err != nil {
// 		log.Fatalln("Exit!", err)
// 	}

// 	log.Println("logging")
// 	log.Printf("%T  %v", "test", "test")

// 	// fatalはexitする
// 	log.Fatalf("%T, %v", "test", "test")
// 	log.Fatalln("error!")
// 	fmt.Println("ok!")
// }

// 28.defer
// func foo() {
// 	defer fmt.Println("world foo")

// 	fmt.Println("hello foo")
// }

// func main() {
// 	// defer fmt.Println("world")

// 	// foo()

// 	// fmt.Println("hello")

// 	// fmt.Println("run")
// 	// defer fmt.Println("1")
// 	// defer fmt.Println("2")
// 	// defer fmt.Println("3")
// 	// fmt.Println("success")

// 	// deferの実用例
// 	file, _ := os.Open("./section4.go")
// 	defer file.Close()
// 	data := make([]byte, 100)
// 	file.Read(data)
// 	fmt.Println(string(data))
// }

// 27.switch
// func getOsName() string {
// 	return "deeefault"
// }

// func main() {
// 	// os := "win"
// 	// 1文でも書ける ※スコープはswitch文の中だけ
// 	switch os := getOsName(); os {
// 	case "mac":
// 		fmt.Println("mac!")
// 	case "windows":
// 		fmt.Println("windows!")
// 	// defaultはなくてもOK
// 	default:
// 		fmt.Println("default!", os)
// 	}

// 	// switchの条件を書かない
// 	t := time.Now()
// 	switch {
// 	case t.Hour() < 12:
// 		fmt.Println("morning!")
// 	case t.Hour() < 17:
// 		fmt.Println("afternoon!")
// 	}
// }

// 26.range
// func main() {
// 	l := []string{"python", "go", "ruby"}
// 	for i := 0; i < len(l); i++ {
// 		fmt.Println(i, l[i])
// 	}

// 	// 違う書き方
// 	for i, v := range l {
// 		fmt.Println(i, v)
// 	}

// 	// iを使用しない
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
// 	for sum < 10 {
// 		sum += sum
// 		fmt.Println(sum)
// 	}
// 	fmt.Println(sum)
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

// 	// num := 11
// 	// if num%2 == 0 {
// 	// 	fmt.Println("by 2")
// 	// } else if num%3 == 0 {
// 	// 	fmt.Println("by 3")
// 	// } else {
// 	// 	fmt.Println("else")
// 	// }

// 	x, y := 11, 12
// 	if x == 10 && y == 10 {
// 		fmt.Println("&&")
// 	}

// 	if x == 10 || y == 10 {
// 		fmt.Println("||")
// 	}
// }
