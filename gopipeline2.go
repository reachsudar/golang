package main

import "fmt"

// main()--> numbergenrator ---> squarer ---> printer---->done

func main() {

	channel1 := make(chan int)
	channel2 := make(chan int)
	channel3 := make(chan struct{})

	go NumberGenerator(channel1, 5)
	go Squarer(channel1, channel2)
	go Printer(channel2, channel3)

	<-channel3

}

func NumberGenerator(in chan int, num int) {
	for i := 0; i < num; i++ {
		in <- i
	}
	close(in)
}
func Squarer(in chan int, out chan int) {
	for a := range in {
		out <- a * a
	}
	close(out)
}
func Printer(in chan int, out chan struct{}) {
	for a := range in {
		fmt.Println(a)
	}
	out <- struct{}{}

}
