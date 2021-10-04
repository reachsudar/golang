package main

import "fmt"

// main()----->number generator ---> printer---->done

func main() {
	//channels
	ch1 := make(chan int)
	done := make(chan struct{})

	// routines
	go number(ch1, 5)
	go printer(ch1, done)
	//done
	<-done
}

//number generator
func number(in1 chan int, num int) {
	for i := 0; i < num; i++ {

		in1 <- i
	}
	close(in1)

}

// printer
func printer(in2 chan int, done chan struct{}) {
	for a := range in2 {
		fmt.Println(a)

	}
	done <- struct{}{}

}
