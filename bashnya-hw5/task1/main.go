package main

import (
	"fmt"
	"sync"
)

func main() {
	numbers := []int{2, 4, 6, 8, 10}

	var wg sync.WaitGroup

	results := make(chan int, len(numbers))

	for _, num := range numbers {
		wg.Add(1)

		go func(n int) {
			defer wg.Done()

			square := n * n
			results <- square
		}(num)
	}

	wg.Wait()

	close(results)

	totalSum := 0
	for res := range results {
		totalSum += res
	}

	fmt.Printf("Последовательность: %v", numbers)
	fmt.Println()
	fmt.Printf("Сумма квадратов: %d", totalSum)
	fmt.Println()
}
