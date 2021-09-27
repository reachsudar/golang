package main

import "fmt"

func main() {
outer:
	for i := 0; i < 3; i++ {
		for j := 0; j < 4; j++ {
			if i == 4 {
				break outer
			} else if i == j {
				continue outer
			}
			fmt.Println(i, " ", j)
		}
	}
}
