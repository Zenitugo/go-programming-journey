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