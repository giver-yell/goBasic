package main

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
