package main

import (
	"fmt"
)

type stack[T any] []T

func (s *stack[T]) Push(t T) {
	*s = append(*s, t)
}

func (s *stack[T]) Pop() {
	*s = (*s)[:len(*s)-1]
}

func (s *stack[T]) Top() T {
	return (*s)[len(*s)-1]
}

func (s *stack[T]) Len() int {
	return len(*s)
}

func (s *stack[T]) isEmpty() bool {
	return s.Len() == 0
}

func main() {
	var s stack[string]
	s.Push("youxuan")

	s.Push("youkai")

	t := s.Top()

	fmt.Println(t)
}
