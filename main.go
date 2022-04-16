package main

import (
	"fmt"
	"strings"
)

const conferenceTickets = 50

var conferenceName = "Ukoo Conference"
var remainingTickets uint = 50
var bookings = []string{}

func main() {

	greetUsers()

	for {
		firstName, lastName, email, userTickets := greetUserInput()
		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets)

		if isValidName && isValidEmail && isValidTicketNumber {

			bookTicket(userTickets, firstName, lastName, email)

			firstNames := getFirstNames()
			fmt.Printf("These are all our bookings first names: %v\n", firstNames)

			if remainingTickets == 0 {
				//end program
				fmt.Println("Tickets are all booked out")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("Your first name or last name is to short, try again")
			}
			if !isValidEmail {
				fmt.Println("Your email formart doesn't contain @ sign, try again")
			}
			if !isValidTicketNumber {
				fmt.Println("The number of tickets is invalid, try again")
			}

		}
	}
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend.")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}

func greetUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter your first name:")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name:")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address:")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets:")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName+" "+lastName+",")

	fmt.Printf("Thank you %v %v for booking %v tickets. You will recieve a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}
