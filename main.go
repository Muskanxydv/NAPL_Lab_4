package main

import (
	"fmt"
	"sync"
)

func computeSquares(n int) int {
	fmt.Printf("Goroutine for computing squares has started\n")
	result := n * n
	fmt.Printf("Goroutine for computing squares has completed\n")
	return result
}

func computeCubes(n int) int {
	fmt.Printf("Goroutine for computing cubes has started\n")
	result := n * n * n
	fmt.Printf("Goroutine for computing cubes has completed\n")
	return result
}

func computeFibonacci(n int) int {
	fmt.Printf("Goroutine for computing Fibonacci has started\n")
	if n <= 1 {
		fmt.Printf("Goroutine for computing Fibonacci has completed\n")
		return n
	}
	result := computeFibonacci(n-1) + computeFibonacci(n-2)
	fmt.Printf("Goroutine for computing Fibonacci has completed\n")
	return result
}

func main() {
	var wg sync.WaitGroup
	var n int
	fmt.Print("Enter a number: ")
	fmt.Scanf("%d", &n)
	results := make(chan int, 3)

	wg.Add(3)
	go func() {
		defer wg.Done()
		results <- computeSquares(n)
	}()
	go func() {
		defer wg.Done()
		results <- computeCubes(n)
	}()
	go func() {
		defer wg.Done()
		results <- computeFibonacci(n)
	}()

	go func() {
		wg.Wait()
		close(results)
	}()

	for result := range results {
		fmt.Println(result)
	}
}
