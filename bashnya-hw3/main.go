package main

import (
	"fmt"

	"bashnya-hw3/stack"
)

func runStackDemo() {
	fmt.Println("--- 1. Демонстрация: Стек ---")
	stack := stack.NewStack()
	stack.Push(100)
	stack.Push(200)
	stack.Push(300)

	fmt.Println("Текущий размер:", stack.Size()) // 3

	item, _ := stack.Pop()
	fmt.Println("Извлекли (Pop):", item)       // 300
	fmt.Println("Новый размер:", stack.Size()) // 2

	stack.Clear()
	fmt.Println("Пустой после Clear()?:", stack.IsEmpty()) // true

	_, err := stack.Pop()
	if err != nil {
		fmt.Println("Ожидаемая ошибка:", err) // stack is empty
	}
}

func main() {
	runStackDemo()

	fmt.Println()
	fmt.Println("---------------------------------")
	fmt.Println()

	runDequeDemo()
}
