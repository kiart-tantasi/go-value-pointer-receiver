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
	// int
	testValue()
	testPointer()
	// string
	testValueString()
	testPointerString()
}

// === functions to test int ==== //
func testValue() {
	var sum = 0
	start := time.Now()
	for i := 0; i < rounds; i++ {
		for j := 0; j < subRounds; j++ {
			sum += computeValue(i, j)
		}
	}
	fmt.Println("\n...int value taken", time.Since(start))
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
	fmt.Println("\n...int pointer taken", time.Since(start))
	fmt.Println("...sum:", sum)
}

// === functions to test string ==== //
const someString = "hello world"

func testValueString() {
	var sum = 0
	start := time.Now()
	for i := 0; i < rounds; i++ {
		for j := 0; j < subRounds; j++ {
			tempString := someString
			sum += computeValueString(tempString)
		}
	}
	fmt.Println("\n...string value taken", time.Since(start))
	fmt.Println("...sum:", sum)
}

func testPointerString() {
	var sum = 0
	start := time.Now()
	for i := 0; i < rounds; i++ {
		for j := 0; j < subRounds; j++ {
			tempString := someString
			sum += computePointerString(&tempString)
		}
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
