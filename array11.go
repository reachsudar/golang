package main

import "fmt"

func main() {

	//Assigning an array

	daysOfWeek := [...]string{
		"mon", "tue", "wed", "thu", "fri",
		"sat", "sun",
	}

	//creating a slice from array

	weekdays := daysOfWeek[0:5]

	//change the value of array

	weekdays[2] = "wednesday"

	// making slice using make()

	listb := make([]int, 2, 5)

	// passing the slice as argument

	printDays(weekdays)

	//passing the array as argument

	printDays(daysOfWeek[:])

	//print lenght and capacity of slice

	printLen(listb)

}

// printDays function

func printDays(days []string) {
	fmt.Printf("%v\n", days)
	fmt.Println()

}

//pint length function

func printLen(days []int) {
	fmt.Printf("%v length:%d  capacity:%d", days, len(days), cap(days))
}
