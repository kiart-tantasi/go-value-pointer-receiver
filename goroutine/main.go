package main

import (
	"fmt"
	"time"
)

const goroutines = 10
const roundsPerGoroutine = 10_000_000_000

func main() {
	testValueGoroutine()
	testPointerGoroutine()
	testValueString()
	testPointerString()
}

// === functions to test int ==== //
func testValueGoroutine() {
	var sum = 0
	start := time.Now()
	values := make(chan int, goroutines)
	for i := 0; i < goroutines; i++ {
		go func(i int) {
			tempSum := 0
			for j := 0; j < roundsPerGoroutine; j++ {
				tempSum += computeValue(i, j)
			}
			values <- tempSum
		}(i)
	}
	for range goroutines {
		sum += <-values
	}
	fmt.Println("\n...int value taken", time.Since(start))
	fmt.Println("...sum:", sum)
}
func testPointerGoroutine() {
	var sum = 0
	start := time.Now()
	values := make(chan int, goroutines)
	for i := 0; i < goroutines; i++ {
		go func(i int) {
			tempSum := 0
			for j := 0; j < roundsPerGoroutine; j++ {
				tempSum += computePointer(&i, &j)
			}
			values <- tempSum
		}(i)
	}
	for range goroutines {
		sum += <-values
	}
	fmt.Println("\n...int pointer taken", time.Since(start))
	fmt.Println("...sum:", sum)
}

// === functions to test string ==== //
const someString = "hello world"

func testValueString() {
	var sum = 0
	start := time.Now()
	values := make(chan int, goroutines)
	for i := 0; i < goroutines; i++ {
		go func(tempString string) {
			tempSum := 0
			for j := 0; j < roundsPerGoroutine; j++ {
				tempSum += computeValueString(tempString)
			}
			values <- tempSum
		}(someString)
	}
	for range goroutines {
		sum += <-values
	}
	fmt.Println("\n...string value taken", time.Since(start))
	fmt.Println("...sum:", sum)
}

func testPointerString() {
	var sum = 0
	start := time.Now()
	values := make(chan int, goroutines)
	for i := 0; i < goroutines; i++ {
		go func(tempString string) {
			tempSum := 0
			for j := 0; j < roundsPerGoroutine; j++ {
				tempSum += computePointerString(&tempString)
			}
			values <- tempSum
		}(someString)
	}
	for range goroutines {
		sum += <-values
	}
	fmt.Println("\n...string pointer taken", time.Since(start))
	fmt.Println("...sum:", sum)
}

// === functions to pretend to do some work === //
func computeValue(i int, j int) int {
	return i + j
}
func computePointer(i *int, j *int) int {
	return *i + *j
}
func computeValueString(str string) int {
	return len(str)
}
func computePointerString(str *string) int {
	return len(*str)
}
