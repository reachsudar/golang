package main

import "fmt"

func main() {

	// Create two channels
	channel1 := make(chan int)
	channel2 := make(chan int)

	//Two go routines
	go squareArea(10, channel1)
	go rectArea(10, 20, channel2)

	//channel receives from worker
	SqArea := <-channel1
	RectArea := <-channel2

	//prints the area
	fmt.Println("Area of square:", SqArea)
	fmt.Println("Area of rectangle:", RectArea)
}

// square worker
func squareArea(side int, sqchannel chan int) {

	area := 4 * side
	sqchannel <- area

}

//rectangle worker
func rectArea(length int, width int, rectchannel chan int) {

	area := length * width
	rectchannel <- area
}
