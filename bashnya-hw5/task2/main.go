package main

import (
	"context"
	"flag"
	"fmt"
	"math/rand"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func worker(ctx context.Context, wg *sync.WaitGroup, id int, dataChan <-chan string) {
	defer wg.Done()
	fmt.Printf("Воркер %d запущен", id)
	fmt.Println()

	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Воркер %d завершает работу (контекст)", id)
			fmt.Println()
			return

		case data, ok := <-dataChan:
			if !ok {
				fmt.Printf("Воркер %d завершает работу (канал закрыт)", id)
				fmt.Println()
				return
			}
			time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
			fmt.Printf("Воркер %d получил: %s", id, data)
			fmt.Println()
		}
	}
}

func main() {
	var numWorkers int
	flag.IntVar(&numWorkers, "n", 3, "Количество воркеров")
	flag.Parse()

	fmt.Printf("Запускаем %d воркеров...", numWorkers)
	fmt.Println()

	dataChan := make(chan string, 10)

	var wg sync.WaitGroup

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(ctx, &wg, i, dataChan)
	}

	go func() {
		i := 0
		for {
			data := fmt.Sprintf("Данные #%d", i)

			select {
			case <-ctx.Done():
				fmt.Println("Продюсер: получен сигнал, завершаю отправку.")
				close(dataChan)
				return

			case dataChan <- data:
				i++
				time.Sleep(100 * time.Millisecond)
			}
		}
	}()

	<-ctx.Done()
	fmt.Println()
	fmt.Println("...Получен сигнал завершения, ожидаем воркеров...")

	wg.Wait()
	fmt.Println("Все воркеры остановлены. Программа завершена.")
}
