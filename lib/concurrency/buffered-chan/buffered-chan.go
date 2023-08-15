package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	fmt.Println("OS: ", runtime.GOOS)                    // GOOS is the running program's operating system target. e.g. darwin
	fmt.Println("ARCH: ", runtime.GOARCH)                // GOARCH is the running program's architecture target. e.g. arm64
	fmt.Println("No. fo CPUs: ", runtime.NumCPU())       // NumCPU returns the number of logical CPUs usable by the current process.
	fmt.Println("Go Routines: ", runtime.NumGoroutine()) // NumGoroutine returns the number of goroutines that currently exist.

	wg := sync.WaitGroup{}
	ch := make(chan int)

	wg.Add(1)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i // Send numbers from 0 to 9
		}
		close(ch)
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			fmt.Println(<-ch) // Prints numbers from 0 to 9
		}
	}()

	fmt.Println("Go Routines: ", runtime.NumGoroutine())

	wg.Wait()

}
