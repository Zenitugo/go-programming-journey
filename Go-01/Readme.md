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