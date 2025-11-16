package main

import (
	"fmt"
	"sync"
)

type ConcurrentMap struct {
	mu   sync.RWMutex
	data map[string]int
}

func NewConcurrentMap() *ConcurrentMap {
	return &ConcurrentMap{
		data: make(map[string]int),
	}
}

func (cm *ConcurrentMap) Set(key string, value int) {
	cm.mu.Lock()
	defer cm.mu.Unlock()
	cm.data[key] = value
}

func (cm *ConcurrentMap) Get(key string) (int, bool) {
	cm.mu.RLock()
	defer cm.mu.RUnlock()
	val, ok := cm.data[key]
	return val, ok
}

func main() {
	cMap := NewConcurrentMap()
	var wg sync.WaitGroup

	numGoroutines := 100
	wg.Add(numGoroutines)

	for i := 0; i < numGoroutines; i++ {
		go func(gIndex int) {
			defer wg.Done()

			key := fmt.Sprintf("key_%d", gIndex%10)
			cMap.Set(key, gIndex)

			// fmt.Printf("Горутина %d записала %s = %d", gIndex, key, gIndex)
			// fmt.Println()
		}(i)
	}

	wg.Wait()

	fmt.Println("Запись завершена. Результат:")
	for i := 0; i < 10; i++ {
		key := fmt.Sprintf("key_%d", i)
		val, ok := cMap.Get(key)
		if ok {
			fmt.Printf("%s = %d", key, val)
			fmt.Println()
		}
	}
}
