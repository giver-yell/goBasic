package main

import (
	"fmt"

	"github.com/giver-yell/goBasic/mylib2"
	"github.com/giver-yell/goBasic/mylib2/under"
)

// 60.package
func main() {
	s := []int{1, 2, 3, 4, 5}
	fmt.Println(mylib2.Average(s))
	mylib2.Say()

	under.Hello()

	// 61.public と private
	person := mylib2.Person{Name: "Mike", Age: 21}
	fmt.Println(person)
}
