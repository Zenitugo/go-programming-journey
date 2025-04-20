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


## Difference between Printf and Println

| Feature              | `fmt.Println()`           | `fmt.Printf()`                            |
|----------------------|---------------------------|-------------------------------------------|
| Adds newline         | Yes                       | No (unless you add `\n`)                  |
| Supports formatting  | No                        | Yes (`%s`, `%d`, `%.2f`, etc.)            |
| Multiple arguments   | Yes (space-separated)     | Yes (with format string)                 |
|


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


## Mini-Project 4: Disk Space Checker
**Description**: Write a program that takes disk space used (in GB, float64) as input and checks if it’s within a safe threshold. If usage is below 50 GB, it’s “Safe”; 50-100 GB is “Warning”; above 100 GB is “Critical”. Print the usage and status.

**Requirements**:

- Prompt: “Enter disk space used in GB: ”
- Input: float64 (e.g., 75.5).
- Output: Usage (1 decimal) and status.


## Mini-Project 5: Backup Frequency Planner
**Description**: Write a program that takes the size of a backup (in MB, int) and suggests a backup frequency. If size is 0-1000 MB, suggest “Daily”; 1001-5000 MB, “Weekly”; above 5000 MB, “Monthly”. Print the size and frequency.

**Requirements**:

- Prompt: “Enter backup size in MB: ”
- Input: int (e.g., 2500).
- Output: Size and frequency.

## Mini-Project 6: Network Latency Monitor
**Description**: Write a program that takes network latency (in milliseconds, float64) and categorizes it. If latency is 0-100 ms, it’s “Excellent”; 101-200 ms, “Acceptable”; above 200 ms, “Poor”. Print latency (2 decimals) and category.

**Requirements**:

- Prompt: “Enter network latency in ms: ”
- Input: float64 (e.g., 150.75).
- Output: Latency (2 decimals) and category.



## Mini-Project 7: CPU Temperature Monitor
**Description**: Write a program that takes a CPU temperature (in Celsius, float64) and categorizes it. If the temperature is 0-60°C, it’s “Normal”; 61-80°C, “Warm”; above 80°C, “Overheating”. Print the temperature (1 decimal) and category.

**Requirements**:

- Prompt: “Enter CPU temperature in Celsius: ”
- Input: float64 (e.g., 72.5).
- Output: Temperature (1 decimal) and category.


## Mini-Project 8: Log Level Classifier
**Description**: Write a program that takes a log level code (integer) and maps it to a severity. Code 1 is “Info”; 2 is “Warning”; 3 is “Error”; anything else is “Unknown”. Print the code and severity.

**Requirements**:

- Prompt: “Enter log level code: ”
- Input: int (e.g., 2).
- Output: Code and severity.




## Mini-Project 9: Memory Allocation Validator
**Description**: Write a program that takes a memory allocation (in MB, int) and categorizes it. If the allocation is 0-512 MB, it’s “Low”; 513-2048 MB, “Moderate”; above 2048 MB, “High”. Print the allocation and category.

**Requirements**:

- Prompt: “Enter memory allocation in MB: ”
- Input: int (e.g., 1024).
- Output: Allocation and category.