First Go Program

package main  // Declares the package name

import "fmt"  // Imports the format package

func main() {  // Main function where execution begins
    fmt.Println("Hello, World!")
}


// Task 1

package main  

import "fmt" 

func main() { 
    var name string = "Ugochi"
    var age int = 35 
    fmt.Printf("Hello, %s! You are %d! years.\n", name, age)
}


// Task 2
package main
import "fmt"

func main() {
    number_of_servers := 10
    server_groups := number_of_servers / 4
    fmt.Printf("The total number of servers is, %d. The number of groups when servers are arrangedin the cluster of 4 is, %d", number_of_servers, server_groups)
}

// Polished version from Grok
package main
import "fmt"
func main() {
    number_of_servers := 10
    server_groups := number_of_servers / 4
    fmt.Printf("The total number of servers is %d. The number of groups when servers are arranged in clusters of 4 is %d.\n", number_of_servers, server_groups)
}