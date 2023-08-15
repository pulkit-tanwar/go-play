package main

import "fmt"

func main() {

	oddChan := make(chan int)
	evenChan := make(chan int)
	quitChan := make(chan int)

	go send(evenChan, oddChan, quitChan)

	receive(evenChan, oddChan, quitChan)
	fmt.Println("End")
}

func send(evenChan, oddChan, quitChan chan<- int) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			evenChan <- i
		} else {
			oddChan <- i
		}
	}
	quitChan <- 0
	close(evenChan)
	close(oddChan)
}

func receive(evenChan, oddChan, quitChan <-chan int) {
	for {
		select {
		case val := <-evenChan:
			fmt.Println("Even chan value: ", val)
		case val := <-oddChan:
			fmt.Println("Odd chan value: ", val)
		case val := <-quitChan:
			fmt.Println("\nquit chan value: ", val)
			return
		}
	}
}
