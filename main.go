package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements generates random elements.
func generateRandomElements(size int) []int {
	// ваш код здесь
	if size <= 0 {
		return []int{}
	}

	s := make([]int, 0, size)
	for i := 0; i < size; i++ {
		s = append(s, rand.Intn(10000))
	}
	return s
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	// ваш код здесь
	if len(data) == 0 {
		return 0
	}

	max := data[0]
	for _, num := range data {
		if num > max {
			max = num
		}
	}
	return max
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	// ваш код здесь

	if len(data) == 0 {
		return 0
	}

	chunkSize := (len(data) + CHUNKS - 1) / CHUNKS
	maxValues := []int{}
	var mu sync.Mutex
	var wg sync.WaitGroup

	chunks := CHUNKS
	if len(data) < CHUNKS {
		chunks = len(data)
	}

	for i := 0; i < chunks; i++ {
		start := i * chunkSize
		end := start + chunkSize
		if start >= len(data) {
			break
		}
		if end > len(data) {
			end = len(data)
		}

		wg.Add(1)
		go func(start, end int) {
			defer wg.Done()
			m := maximum(data[start:end])
			mu.Lock()
			maxValues = append(maxValues, m)
			mu.Unlock()
		}(start, end)
	}
	wg.Wait()

	if len(maxValues) == 0 {
		return 0
	}

	return maximum(maxValues)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел", SIZE)
	// ваш код здесь
	data := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	// ваш код здесь
	start := time.Now()
	max := maximum(data)
	elapsed := time.Since(start).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d mcs\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	// ваш код здесь
	start = time.Now()
	max = maxChunks(data)
	elapsed = time.Since(start).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d mcs\n", max, elapsed)
}
