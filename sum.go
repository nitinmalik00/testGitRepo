package main

import "fmt"

func sum(x, y int) int {
	return x + y
}

func main() {
	fmt.Println("Hello World, this is for summing two numbers")

	result := sum(5, 7)
	fmt.Println("The sum is: ", result)
}
