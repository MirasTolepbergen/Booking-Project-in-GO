package main

import (
	"fmt"
	"sync"
	"time"
)

const conferenceTickets int = 50

var conferenceName string = "Red Pill Conference"
var remainingTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	greetUsers()
	//to keep asking for user input

	firstName, lastName, email, userTickets := getUserInput()
	isValidName, isValidEmail, isValidTicketNumber := ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

	if isValidName && isValidEmail && isValidTicketNumber {

		bookTicket(userTickets, firstName, lastName, email)

		wg.Add(1)
		//concurency
		go sendTicket(userTickets, firstName, lastName, email)

		//to display only first names
		firstNames := getFirstNames()
		fmt.Printf("these are all the bookings: %v\n", firstNames)

		if remainingTickets == 0 {
			//quit if no tickets left
			fmt.Println("Our conference is booked out.")
		}
	} else {
		if !isValidName {
			fmt.Println("first name/last name is too short")
		}
		if !isValidTicketNumber {
			fmt.Println("number of tickets your entered is invalid")
		}
		if !isValidEmail {
			fmt.Println("email your entered is invalid")
		}
	}
	wg.Wait()
}

func greetUsers() {
	fmt.Printf("welcome to %v booking application", conferenceName)
	fmt.Printf("we have total of %v tickets and %v still remain.\n", conferenceTickets, remainingTickets)
	fmt.Println("get ur tickets please!")
}

func getFirstNames() []string {
	firstNames := []string{}

	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames

}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter ur 1st name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter ur last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter ur email: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	//create a map for a user
	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thanks %v %v for booking %v tickets. U'll receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(20 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Printf("--------------------------------------\n")
	fmt.Printf("Sending ticket:\n %v to email adress %v\n", ticket, email)
	fmt.Printf("--------------------------------------\n")
	wg.Done()
}
