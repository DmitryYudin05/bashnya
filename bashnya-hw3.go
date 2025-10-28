package main

import (
	"fmt"
)

type Stack struct {
	items []interface{}
}

func (s *Stack) Push(value interface{}) {
	s.items = append(s.items, value)
}
func (s *Stack) Pop() (interface{}, bool) {
	if s.IsEmpty() {
		return nil, false
	}
	index := len(s.items) - 1
	element := s.items[index]
	s.items = s.items[:index]
	return element, true
}
func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}
func (s *Stack) Size() int {
	return len(s.items)
}
func (s *Stack) Clear() {
	s.items = nil
}
func (s *Stack) VivodVStroku() string {
	return fmt.Sprintf("%v", s.items)
}
func main() {
	stack := &Stack{}
	chislo := 0
	elem := 0
	fmt.Println("Стек изначально:", stack.VivodVStroku())
	fmt.Println("Пожалуйста введите число элементов стека которые хотите добавить:")
	fmt.Scan(&chislo)
	for i := 1; i <= chislo; i++ {
		fmt.Println("Пожалуйста введите элемент:")
		fmt.Scan(&elem)
		stack.Push(elem)
	}
	fmt.Println("Стек после добавления:", stack.VivodVStroku())
	item, ok := stack.Pop()
	if ok {
		fmt.Printf("Удаляем последний элемент %v\n", item)
	}
	fmt.Println("Стек после удаления:", stack.VivodVStroku())
	fmt.Println("Размер стека:", stack.Size())
	stack.Clear()
	fmt.Println("Стек после очистки:", stack.VivodVStroku())
	fmt.Println("Пустой ли стек?", stack.IsEmpty())
}
