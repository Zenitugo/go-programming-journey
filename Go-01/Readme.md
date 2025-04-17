# KEY CONCEPTS

## Program Structure

#### Every Go program needs a package main and a func main()—it’s where execution starts.

```
package main
import "fmt"
func main() {
    fmt.Println("Hello, World!")
}

```
- fmt is a package for printing (like Python’s print()).


## Variables
- Declare with var or := (short form, only inside functions).
- You must specify types (e.g., string, int, float64).

```
var name string = "Amaka"
age := 25  // Short form, infers int

```

```

# Explicit declaration
var name string = "John"

# Type inference
var age = 30
height := 175.5  # Short declaration (only inside functions)

# Multiple variables
var (
    a int = 1
    b     = 2
    c     = 3
)

 #Constants
const pi = 3.14

```

## Output
- Use fmt.Println() for output.
- Use fmt.Printf() for formatted output (e.g., fmt.Printf("Name: %s", name)).
- %s for strings, %d for integers, %f for floats.

## Input:
- Go’s input is trickier than Python’s input(). We’ll use fmt.Scan() or fmt.Scanln() for now.

```
var input string
fmt.Scan(&input)

```

## Basic Data Types
- bool: true/false

- string: "text"

- int, int8, int16, int32, int64: integers

- uint, uint8, uint16, uint32, uint64: unsigned integers

- float32, float64: floating-point numbers

- complex64, complex128: complex numbers

# Tasks

## Task 1: Hello Greeter

**Write a program that:**
- Declares a variable for your name (string).
- Declares a variable for your age (integer).
- Prints: "Hello, [name]! You are [age] years


## Task 2: Server Count
**Write a program that:**
- Declares a variable for the number of servers (integer).
- Calculates how many can be grouped into clusters of 4 (use /).
- Prints both numbers.

## Task 3: Uptime Printer
**Write a program that:**
- Takes user input for uptime in hours (integer, use fmt.Scan()).
- Prints it with a message


## Mini-Project 1: Memory Converter
Ask for memory in MB (float64, use fmt.Scan()).
Convert to GB (divide by 1024).
Print both, rounded to 1 decimal place (use %.1f).


## Mini-Project 2: Status Checker
- Ask for a server status code (integer, use fmt.Scan()).
- Map it to a message:
1: "Server online".
2: "Server offline".
Else: "Unknown status".
Print the message.


## Mini-Project 3: Port Validator
1. Ask for a port number (integer, use fmt.Scan()).
2. Check if it’s valid (1-65535).
3. Print "Valid port" or "Invalid port".