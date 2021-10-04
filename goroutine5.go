package main

import "fmt"

func main() {
	channel1 := make(chan int)

	go count(channel1)

	// Recieve signal line by line
	/*a := <-channel1
	fmt.Println(a)

	a = <-channel1
	fmt.Println(a)*/
	for a := range channel1 {
		fmt.Println(a)
	}

}
func count(recv chan int) {
	for i := 0; i < 5; i++ {
		recv <- i
	}
	close(recv)

}
