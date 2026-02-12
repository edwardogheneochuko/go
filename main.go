
package main

import "fmt"

func main() {	
	var decideName = "ride Decide"
	const decideTickets = 40
	var ridingTickets = 25

	fmt.Printf("Welcome to the %v party\n", decideName)
	fmt.Printf("A total of tickets %v tickets which are riding %v Tickets are still avaliable.\n", decideTickets, ridingTickets )
	fmt.Println("Let those who", decideName," ...!!!")
}