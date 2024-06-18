package main

import (
	"fmt"
	"time"
)

const goroutines = 10
const roundsPerGoroutine = 10_000_000_000

func main() {
	testValueChannel()
	testPointerChannel()
}
func testValueChannel() {
	sum := 0
	start := time.Now()
	in := make(chan int, 10)
	out := make(chan int)
	for i := range goroutines {
		in <- i
	}
	for range goroutines {
		go sumValue(in, out)
	}
	close(in)
	for range goroutines {
		sum += <-out
	}
	fmt.Printf("Sum by value: %d (%v)\n", sum, time.Since(start))
}
func sumValue(in chan int, out chan int) {
	i := <-in
	tempSum := 0
	for j := range roundsPerGoroutine {
		tempSum += computeValue(i, j)
	}
	out <- tempSum
}

func testPointerChannel() {
	sum := 0
	start := time.Now()
	in := make(chan int, 10)
	out := make(chan int)
	for i := range goroutines {
		in <- i
	}
	for range goroutines {
		go sumPointer(in, out)
	}
	close(in)
	for range goroutines {
		sum += <-out
	}
	fmt.Printf("Sum by pointer: %d (%v)\n", sum, time.Since(start))
}
func sumPointer(in chan int, out chan int) {
	i := <-in
	tempSum := 0
	for j := range roundsPerGoroutine {
		tempSum += computePointer(&i, &j)
	}
	out <- tempSum
}

// === functions to pretend to do some work === //
func computeValue(i int, j int) int {
	return i + j
}
func computePointer(i *int, j *int) int {
	return *i + *j
}
