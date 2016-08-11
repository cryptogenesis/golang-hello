package main

import "fmt"

const (
msg1 = "constant -> hello world"
)

func main() {
	msg := "Hello world in golang!"
	fmt.Println("hello", "world", "in", "golang!")
	fmt.Println(msg)
	fmt.Println(msg1)
}
