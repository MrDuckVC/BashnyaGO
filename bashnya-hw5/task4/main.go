package main

import (
	"fmt"
	"sync"
)

func generator(wg *sync.WaitGroup, nums []int, out chan<- int) {
	defer wg.Done()
	for _, n := range nums {
		out <- n
	}
	close(out)
}

func multiplier(wg *sync.WaitGroup, in <-chan int, out chan<- int) {
	defer wg.Done()
	for n := range in {
		out <- (n * 2)
	}
	close(out)
}

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	ch1 := make(chan int)
	ch2 := make(chan int)

	var wg sync.WaitGroup
	wg.Add(2)

	go generator(&wg, numbers, ch1)
	go multiplier(&wg, ch1, ch2)

	fmt.Println("Результаты конвейера (x * 2):")
	for result := range ch2 {
		fmt.Println(result)
	}

	wg.Wait()
	fmt.Println("Конвейер завершил работу.")
}
