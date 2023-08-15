package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3
	<-ch // Here 1 will come
	close(ch)
	a := <-ch         // here 2 will come
	b := <-ch         // Now here the Default 0 value of int will come.
	fmt.Println(a, b) // Ans : 2 0

	val, ok := <-ch
	fmt.Println("ok: ", ok)
	fmt.Println("ok: ", val)

}

// If We can receive value from a closed channel, it will not give an error.
// If there are values in the buffer, we will get them; otherwise, we will get the zero value for the channel type.

// But we cannot send values to a closed channel.

// We can use the ,ok (comma ok notation) to check if the 0 value received from the channel was from the closed channel
