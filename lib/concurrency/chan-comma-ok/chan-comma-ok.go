package main

import "fmt"

func main() {

	ch := make(chan int, 3)

	ch <- 2
	ch <- 9
	ch <- 4

	fmt.Println(<-ch) // 2

	val, ok := <-ch
	fmt.Printf("%t, %d\n", ok, val) // true, 9

	fmt.Println(<-ch)

	fmt.Println(<-ch) // Deadlock

}
