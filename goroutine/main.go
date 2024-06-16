package main

import (
	"fmt"
	"sync"
	"time"
)

const goroutines = 10
const roundsPerGoroutine = 10_000_000

var mutex sync.Mutex

func main() {
	testValueGoroutine()
	testPointerGoroutine()
}

// === functions to test ==== //
func testValueGoroutine() {
	var sum = 0
	var wg sync.WaitGroup
	start := time.Now()
	for i := 0; i < goroutines; i++ {
		wg.Add(1)
		go func(i int) {
			tempSum := 0
			for j := 0; j < roundsPerGoroutine; j++ {
				tempSum += computeValue(i, j)
			}
			mutex.Lock()
			sum += tempSum
			mutex.Unlock()
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println("\n...value (goroutine) taken", time.Since(start))
	fmt.Println("...sum:", sum)
}
func testPointerGoroutine() {
	var sum = 0
	var wg sync.WaitGroup
	start := time.Now()
	for i := 0; i < goroutines; i++ {
		wg.Add(1)
		go func(i int) {
			tempSum := 0
			for j := 0; j < roundsPerGoroutine; j++ {
				tempSum += computePointer(&i, &j)
			}
			mutex.Lock()
			sum += tempSum
			mutex.Unlock()
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println("\n...pointer (goroutine) taken", time.Since(start))
	fmt.Println("...sum:", sum)
}

// === functions to pretend to do some work === //
func computeValue(i int, j int) int {
	return i * j
}
func computePointer(i *int, j *int) int {
	return *i * *j
}
