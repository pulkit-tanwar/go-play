package main

import (
	"fmt"
	"sync"
)

func main() {

	wg := sync.WaitGroup{}
	ch := make(chan int)

	wg.Add(1)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
	}()

	go func() {
		defer wg.Done()

		for val := range ch { // To use range, we need to close the channel.
			fmt.Println(val)
		}

	}()

	wg.Wait()

}
