package main

import (
	"fmt"
	"github.com/golang-collections/collections/stack"
)

func main() {
	s := stack.New()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Pop()
	fmt.Println(s.Peek())
	fmt.Println("Hello, World!")

}
