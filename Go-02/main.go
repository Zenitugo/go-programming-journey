package main

import "fmt"

func main() {
	var n int
	fmt.Print("Enter a positive number: ")
	fmt.Scan(&n)

	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}

	fmt.Printf("Sum: %d\n", sum)
}

