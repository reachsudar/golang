package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Begins")
	go delayedPrint("First")
	go delayedPrint("Second")
	go delayedPrint("Third")

	fmt.Println("Ends")
	time.Sleep(3 * time.Second)
}

func delayedPrint(S string) {
	for i := 0; i < 5; i++ {
		fmt.Printf("%d %s", i, S)
		fmt.Println()
		time.Sleep(1 * time.Second)
	}
}
