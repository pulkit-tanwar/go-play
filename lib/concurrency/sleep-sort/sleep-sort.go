package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	sleepSort1() // Prints 2 2 2 2 2
	sleepSort2() // Prints 1 2 3 4 5
	sleepSort3() // Prints 1 2 3 4 5
}

func sleepSort1() {
	wg := sync.WaitGroup{}
	slice := []int{4, 5, 3, 1, 2}

	fmt.Printf("sleepSort1: ")
	for _, n := range slice {
		wg.Add(1)
		go func() {
			defer wg.Done()
			time.Sleep(time.Duration(n) * time.Millisecond)
			fmt.Printf("%d ", n) // Prints 2 2 2 2 2 Why ? Because of loop closure
		}()
	}

	wg.Wait()
	fmt.Println()
}

func sleepSort2() {
	wg := sync.WaitGroup{}
	slice := []int{4, 5, 3, 1, 2}

	fmt.Printf("sleepSort2: ")
	for _, n := range slice {
		wg.Add(1)
		n := n // Declaring a new variable
		go func() {
			defer wg.Done()
			time.Sleep(time.Duration(n) * time.Millisecond)
			fmt.Printf("%d ", n) // Prints 1 2 3 4 5
		}()
	}
	wg.Wait()
	fmt.Println()
}

func sleepSort3() {
	wg := sync.WaitGroup{}
	slice := []int{4, 5, 3, 1, 2}

	fmt.Printf("sleepSort3: ")
	for _, n := range slice {
		wg.Add(1) // Passing argument to func
		go func(i int) {
			defer wg.Done()
			time.Sleep(time.Duration(i) * time.Millisecond)
			fmt.Printf("%d ", i) // Prints 1 2 3 4 5
		}(n)
	}
	wg.Wait()
	fmt.Println()
}
