package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Begins")
	delayedPrint()
	fmt.Println("Ends")
}

func delayedPrint() {

	time.Sleep(2 * time.Second)
	fmt.Println("Lazy print")

}
