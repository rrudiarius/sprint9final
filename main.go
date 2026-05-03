package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements generates random elements.
func generateRandomElements(size int) []int {
	if size <= 0 {
		return nil
	}

	num := make([]int, 0, size)

	rnd := rand.NewSource(time.Now().Unix())
	r := rand.New(rnd)

	for i := 0; i < size; i++ {
		num = append(num, r.Int())
	}

	return num
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	if len(data) == 0 {
		return 0
	}

	max := data[0]
	for _, v := range data {
		if v > max {
			max = v
		}
	}

	return max
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	if len(data) == 0 {
		return 0
	}

	sliceSize := len(data) / CHUNKS
	remainder := len(data) % CHUNKS

	maxSlice := make([]int, CHUNKS)

	start := 0
	for i := 0; i < CHUNKS; i++ {
		// определяем размер текущего чанка
		end := start + sliceSize
		if i < remainder {
			end++ // добавляем один элемент из остатка
		}

		partSlice := data[start:end]

		wg.Add(1)

		go func(m int, p []int) {
			defer wg.Done()
			maxSlice[m] = maximum(p)
		}(i, partSlice)

		start = end
	}

	wg.Wait()

	return maximum(maxSlice)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	generation := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	timeGen := time.Now()
	max := maximum(generation)
	elapsed := time.Since(timeGen)
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d мкс\n", max, elapsed.Microseconds())

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	timeGen = time.Now()
	max = maxChunks(generation)
	elapsed = time.Since(timeGen)
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d мкс\n", max, elapsed.Microseconds())
}
