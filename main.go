package main

import(
	"fmt"
	"strings"
)

func main() {
	var conferenceName string = "Go conference"  // or conferenceName := "Go conference"	
	const conferenceTickets int = 50
	var remainigTickets uint = 50
	var bookings = []string{}   // or bookings := []string{}
	//var bookings [50]string  

	fmt.Printf("conferenceName is %T, conferenceTickets is %T, remainigTickets is %T\n", conferenceName, conferenceTickets, remainigTickets )

	fmt.Println("Welcome to", conferenceName, "booking application")
	// fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Println("We have total of", conferenceTickets, "tickets and", remainigTickets, "are still available." )
	fmt.Println("Get your tickets here to attend")
	
	for {
		var firstName string
		var lastName string
		var email string
		var userTickets uint

		fmt.Println("Enter your first name: ")
		fmt.Scan(&firstName)
		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName)
		fmt.Println("Enter your email address: ")
		fmt.Scan(&email)
		fmt.Println("Enter number of tickets: ")
		fmt.Scan(&userTickets)

		remainigTickets = remainigTickets - userTickets
		//bookings[0] = firstName + " " + lastName
		bookings = append(bookings, firstName + " " + lastName)

		fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
		fmt.Printf("%v tickets remaining for %v\n", remainigTickets, conferenceName)	
		
		firstNames := []string{}
		for _, b := range bookings {
			var names = strings.Fields(b)
			firstNames = append(firstNames, names[0] )
		}
		fmt.Printf("The first names of bookings are: %v\n", firstNames)
	}
	
	

	

	
}