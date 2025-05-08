// Mini Project 1
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

// Grok's Solution to Mini Project 1	
package main

import (
	"fmt"
)
func sumNumbers(num int) int {
	total := 0
	for i := 1; i <= num; i++ {
		total += i
	}
	return total
}

func main() {
	var num int
	fmt.Print("Enter a positive number: ")
	_, err := fmt.Scanln(&num)
	if err != nil || num <= 0 {
		fmt.Println("Invalid input: Enter a positive number")
		return
	}
	result := sumNumbers(num)
	fmt.Printf("Sum: %d\n", result)
}