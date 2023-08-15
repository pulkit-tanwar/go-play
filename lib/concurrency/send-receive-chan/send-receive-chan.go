package main

import (
	"fmt"
)

func main() {

	c := make(chan int)

	go send(c)
	receive(c)

	// time.Sleep(time.Second * 3)
	fmt.Println("Program Ended!")
}

func send(c chan<- int) {
	c <- 30
}

func receive(c <-chan int) {
	ans := <-c
	fmt.Println(ans)
}
