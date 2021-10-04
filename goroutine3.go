package main

import "fmt"

func main() {
	done := make(chan bool)
	fmt.Println("Begins")
	go Function1(done)
	<-done
	fmt.Println("Ends")

}
func Function1(here chan bool) {

	fmt.Println("Inside Function")
	here <- true

}
