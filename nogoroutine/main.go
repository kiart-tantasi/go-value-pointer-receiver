package main

import (
	"fmt"
	"sync"
	"time"
)

const rounds = 10
const subRounds = 10_000_000

var mutex sync.Mutex

func main() {
	testValue()
	testPointer()
}

// === functions to test ==== //
func testValue() {
	var sum = 0
	start := time.Now()
	for i := 0; i < rounds; i++ {
		for j := 0; j < subRounds; j++ {
			sum += computeValue(i, j)
		}
	}
	fmt.Println("\n...value taken", time.Since(start))
	fmt.Println("...sum:", sum)
}
func testPointer() {
	var sum = 0
	start := time.Now()
	for i := 0; i < rounds; i++ {
		for j := 0; j < subRounds; j++ {
			sum += computePointer(&i, &j)
		}
	}
	fmt.Println("\n...pointer taken", time.Since(start))
	fmt.Println("...sum:", sum)
}

// === functions to pretend to do some work === //
func computeValue(i int, j int) int {
	return i * j
}
func computePointer(i *int, j *int) int {
	return *i * *j
}
