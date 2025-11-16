package main

import "fmt"

func setBit(n int64, pos uint) int64 {
	n |= (1 << pos)
	return n
}

func clearBit(n int64, pos uint) int64 {
	n &= ^(1 << pos)
	return n
}

func setBitValue(n int64, pos uint, value bool) int64 {
	if value {
		return setBit(n, pos)
	}
	return clearBit(n, pos)
}

func main() {
	var n int64 = 10

	fmt.Printf("Исходное число: %d (%b)", n, n)
	fmt.Println()

	// Устанавливаем 0-й бит (pos=0) в 1
	n = setBitValue(n, 0, true)
	fmt.Printf("Установить бит 0: %d (%b)", n, n)
	fmt.Println()

	// Устанавливаем 1-й бит (pos=1) в 0
	n = setBitValue(n, 1, false)
	fmt.Printf("Сбросить бит 1:   %d (%b)", n, n)
	fmt.Println()

	// Устанавливаем 3-й бит (pos=3) в 0
	n = setBitValue(n, 3, false)
	fmt.Printf("Сбросить бит 3:   %d (%b)", n, n)
	fmt.Println()

	// Устанавливаем 4-й бит (pos=4) в 1
	n = setBitValue(n, 4, true)
	fmt.Printf("Установить бит 4: %d (%b)", n, n)
	fmt.Println()
}
