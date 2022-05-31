package main

// import (
// 	"fmt"

// "github.com/giver-yell/goBasic/mylib2"
// "github.com/giver-yell/goBasic/mylib2/under"
// )

import (
	"fmt"

	"github.com/markcheno/go-quote"
	"github.com/markcheno/go-talib"
)

// 64.サードパーティのpackageのインストール
func main() {
	spy, _ := quote.NewQuoteFromYahoo("spy", "2016-01-01", "2016-04-01", quote.Daily, true)
	fmt.Print(spy.CSV())
	rsi2 := talib.Rsi(spy.Close, 2)
	fmt.Println(rsi2)
}

// 60.package
// func main() {
// 	s := []int{1, 2, 3, 4, 5}
// 	fmt.Println(mylib2.Average(s))
// 	mylib2.Say()

// 	under.Hello()

// 	// 61.public と private
// 	person := mylib2.Person{Name: "Mike", Age: 21}
// 	fmt.Println(person)
// }
