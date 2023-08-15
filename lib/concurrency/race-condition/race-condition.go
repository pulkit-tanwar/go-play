package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	race1() // wrong output, has race conditions
	race2() // solving using Mutex
	race3() // solving using Atomic Package
}

func race1() {
	var count int
	var wg sync.WaitGroup

	for i := 0; i < 1_000_000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			count++
		}()
	}
	wg.Wait()
	fmt.Println("Count from race1(): ", count) // It will be a large number less than 1000000
}

func race2() { // Solving using Mutex
	wg := sync.WaitGroup{}
	var count int = 0

	mutex := sync.Mutex{}

	for i := 0; i < 1_000_000; i++ {
		wg.Add(1)
		go func() {
			mutex.Lock() // Acquiring the lock
			defer wg.Done()
			defer mutex.Unlock()
			count++
		}()
	}
	wg.Wait()
	fmt.Println("Count from race2(): ", count) // Prints 1000000
}

func race3() { // Solving using Atomic
	wg := sync.WaitGroup{}
	var count int64

	for i := 0; i < 1_000_000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomic.AddInt64(&count, 1)

		}()
	}
	wg.Wait()
	fmt.Println("Count from race3(): ", count) // Prints 1000000
}
