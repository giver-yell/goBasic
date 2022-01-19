package mylib

import "fmt"

// public は Person（大文字）
type Person struct {
	Name string
	Age  int
}

func Say() {
	fmt.Println("Human!")
}
