package main

import (
	"booking-app/helpers"
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

var conferenceName = "Go Conference"

const conferenceTickets = 50

var remainingTickets = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets int
}

func main() {

	greetUser(conferenceName, conferenceTickets, remainingTickets)

	for {

		firstName, lastName, email, userTickets := getInput()

		isValidName, isValidEmail, isValidUserTickets := helpers.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidUserTickets {

			bookTickets(userTickets, firstName, lastName, email)
			wg.Add(1)
			go sendTicket(userTickets, firstName, lastName, email)

			fmt.Printf("These are all our bookings %v\n", makeFirstNames())

		} else {
			if !isValidName {
				fmt.Printf("First Name or Last Name is too short...")
				continue
			}
			if !isValidEmail {
				fmt.Printf("Email format is not right")
				continue
			}
			if !isValidUserTickets {
				fmt.Printf("Number of tickets you entered is invalid")
			}

		}

		noRemainingTickets := remainingTickets == 0
		if noRemainingTickets {
			fmt.Printf("Our conference is booked out. Come back next year.")
			break
		} else {
			continue
		}
	}
	wg.Wait()
}

func greetUser(greetName string, greetTickets int, greetRemaining int) {

	fmt.Printf("Welcome to our %v booking application\n", greetName)
	fmt.Printf("Tickets are: %v$, Remaining tickets: %v\n", greetTickets, greetRemaining)
	fmt.Println("Get your tickets here to attend")
}

func makeFirstNames() []string {

	var firstNames []string
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames

}

func getInput() (string, string, string, int) {

	var firstName string
	var lastName string
	var email string
	var userTickets int

	fmt.Print("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Print("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Print("Enter your email: ")
	fmt.Scan(&email)

	fmt.Print("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets

}

func bookTickets(userTickets int, firstName string, lastName string, email string) {

	remainingTickets -= userTickets

	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will recieve a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

}

func sendTicket(userTickets int, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("*******************")
	fmt.Printf("Sending ticket: \n %v \nto email address %v\n", ticket, email)
	fmt.Println("*******************")
	wg.Done()
}
