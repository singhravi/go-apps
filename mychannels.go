package main

import "fmt"



func mychannels(){
	unbuffered := make(chan int)
	buffered := make(chan string, 10)
	//sending data inside channel is done using <-
	buffered <- "Ravi"
	unbuffered <- 100
	//For another gorou- tine to receive that string from the channel, 
	//we use the same <- operator, but this time as a unary operator.
	value := <- buffered
	intValue := <-unbuffered
	fmt.Println(value)
	fmt.Println(intValue)

}