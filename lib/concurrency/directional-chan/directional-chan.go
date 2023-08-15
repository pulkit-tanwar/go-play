package main

import "fmt"

func main() {
	c1 := make(chan int, 1) // Bi Directional Channel

	c1 <- 20

	fmt.Println(<-c1)
	fmt.Printf("Type of c1 : %T \n", c1) // Type - chan int

	cSend := make(chan<- int, 1) // send-only channel (You can only send values to this channel)
	cSend <- 30

	fmt.Printf("Type of ch : %T (send-only channel)\n", cSend) // Type - chan<- int (send-only channel)

	cReceive := make(<-chan int, 1)
	fmt.Printf("Type of ch : %T (receive-only channel)\n", cReceive) // Type - <-chan int (receive-only channel)

	// fmt.Println(<-cSend) // Error : invalid operation: cannot receive from send-only channel cSend (variable of type chan<- int)
	// cReceive <- 40 // Error : invalid operation: cannot send to receive-only channel cReceive (variable of type <-chan int)

	cReceive = c1 // works
	cSend = c1    // works

	// c1 = cReceive // annot use cReceive (variable of type <-chan int) as chan int value in assignment
	// c1 = cSend // cannot use cSend (variable of type chan<- int) as chan int value in assignment

}
