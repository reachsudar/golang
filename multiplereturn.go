package main

import "fmt"

func rectArea(length float64, width float64) (float64, float64) {
	var area = length * width
	var perimeter = (length + width) * 2
	return area, perimeter
}
func main() {
	area, perimeter := rectArea(10.8, 9.5)
	fmt.Println("Area=", area, "perimeter=", perimeter)
}
