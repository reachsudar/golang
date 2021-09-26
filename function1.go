package main

import "fmt"

func main() {

	message("john", 33)
	message("sam", 35)
	fmt.Println(result(3, 6))
}
func result(x int, y int) (z int) {
	z = x + y
	return
}

func message(name string, age int) {
	fmt.Println("I am", name, "Age", age)
}
