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


// Polished Version froom Grok for mini project 3

package main
import "fmt"
func main() {
    var port int
    fmt.Print("Enter port number: ")
    fmt.Scan(&port)
    if port >= 1 && port <= 65535 {
        fmt.Println("Valid port")
    } else {
        fmt.Println("Invalid port")
    }
}

// Mini Project 4

package main

import "fmt"

func main() {
    var disk_space_used float64

    fmt.Print("What is the disk space used in GB? ")
    fmt.Scan(&disk_space_used)

    var status string 

    if disk_space_used < 50{
        status = "Safe"
    } else if disk_space_used >= 50 && disk_space_used <= 100 {
        status = "Warning"
    } else {
        status = "Critical"
    }
    fmt.Printf("Disk space usage is: %.1f GB, Status is: %s\n", disk_space_used, status)
}


// Mini Project 5
package main

import "fmt"

func main(){
    var backup_size int

    fmt.Print("Enter the backup size in MB: ")
    fmt.Scan(&backup_size)

    var backup_frequency string

    if backup_size >= 0 && backup_size <= 1000 {
        backup_frequency = "Daily"
    } else if backup_size >= 1001 && backup_size <= 5000 {
        backup_frequency = "Weekly"
    } else {
        backup_frequency = "Monthly"
    }

    fmt.Printf("Backup size: %d MB, Frequency: %s\n", backup_size, backup_frequency)
}


// Mini Project 6

package main

import "fmt"

func main() {
    var network_latency float64

    fmt.Print("Enter network latency in milliseconds: ")
    fmt.Scan(&network_latency)

    var status string

    if network_latency >=0 && network_latency <= 100 {
        status = "Excellent"
    } else if network_latency >= 101 && network_latency <= 200 {
        status = "Acceptable"
    } else {
        status = "Poor"
    }

    fmt.Printf("Network latency: %.2f milliseconds, Category: %s\n", network_latency, status)
}

// Polished version of mini project 6 from Grok
package main
import "fmt"
func main() {
    var network_latency float64
    fmt.Print("Enter network latency in ms: ")
    fmt.Scan(&network_latency)
    var category string
    if network_latency >= 0 && network_latency <= 100 {
        category = "Excellent"
    } else if network_latency <= 200 {
        category = "Acceptable"
    } else {
        category = "Poor"
    }
    fmt.Printf("Latency: %.2f ms, Category: %s\n", network_latency, category)
}



// Mini Project 7
package main

import "fmt"

func main() {
	var cpu_temp float64
	fmt.Print("Enter the CPU temperature in celsius: ")
	fmt.Scan(&cpu_temp)

	var status string
	if cpu_temp >= 0 && cpu_temp <= 60 {
		status = "Normal"
	} else if cpu_temp <= 80 {
		status = "Warm"
	} else {
		status = "Overheating"
	}

	fmt.Printf("The CPU temperature is %.1f °C, It is currently %s\n", cpu_temp, status)
}