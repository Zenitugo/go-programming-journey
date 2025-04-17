// First Go Program

package main  // Declares the package name

import "fmt"  // Imports the format package

func main() {  // Main function where execution begins
    fmt.Println("Hello, World!")
}


// Task 1 - Hello Greeter

package main  

import "fmt" 

func main() { 
    var name string = "Ugochi"
    var age int = 35 
    fmt.Printf("Hello, %s! You are %d! years.\n", name, age)
}


// Task 2 - Server Count
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



// Task 3 - Uptime Printer

package main
import "fmt"

func main() {
    
    var uptime_hours int input
    fmt.Scan("What is the uptime in hours? ")
    fmt.Printf("The uptime in hours is %d," uptime_hours )
}


// Polished version from Grok

package main
import "fmt"
func main() {
    var uptime_hours int
    fmt.Print("Enter uptime in hours: ")
    fmt.Scan(&uptime_hours)
    fmt.Printf("Server uptime: %d hours\n", uptime_hours)
}



// Mini Project 1
package main

import "fmt"

func main() {

var memory float64
fmt.Print("What is the compute memory in MB? ")
fmt.Scan(&memory)

quotient := memory / 1024
fmt.Printf("The memory in GB is: %.1f\n", quotient)

}


// Mini Project 2

package main
import "fmt"

func main() {
    var status int
    fmt.Print("What is the server status code? ")
    fmt.Scan(&status)

    if status == 1{
        fmt.Printf("Server Online\n")
    } else if status == 2 {
        fmt.Printf("Server Offline\n")
    }else {
        fmt.Printf("Unknown Status\n")
    }
}


// Mini Project 3

package main

import "fmt"

func main() {
    var port int
    
    fmt.Print("What is the port number? ")
    fmt.Scan(&port)

    if port >=1 && port <= 65535{
        fmt.Printf("Valid port\n")
    } else {
        fmt.Printf("Invalid port\n")
    }
}