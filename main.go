package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Ukoo Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50
	bookings := []string{}

	fmt.Printf("Welcome to %v booking application.\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend.")

	for {
		var firstName string
		var lastName string
		var email string
		var userTickets uint
		// ask user name
		fmt.Println("Enter your first name:")
		fmt.Scan(&firstName)

		fmt.Println("Enter your last name:")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email address:")
		fmt.Scan(&email)

		fmt.Println("Enter number of tickets:")
		fmt.Scan(&userTickets)

		if userTickets <= remainingTickets {
			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, firstName+" "+lastName+",")

			fmt.Printf("Thank you %v %v for booking %v tickets. You will recieve a confirmation email at %v\n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceTickets)

			//Display only first name
			// This loop ends when iterated over all elements of te booking list
			firstNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0]+",")
			}

			fmt.Printf("These are all our bookings: %v\n", bookings)
			fmt.Printf("These are all our bookings first names: %v\n", firstNames)

			if remainingTickets == 0 {
				//end program
				fmt.Println("Tickets are all booked out")
				break
			}
		} else {
			fmt.Printf("You cannot book %v tickets. You can only book from the remaining %v tickets\n", userTickets, remainingTickets)
		}
	}
}
