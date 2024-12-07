package main

import (
	"booking-app/utils"
	"fmt"
	"sync"
	"time"
)

const conferenceTickets uint = 50

var conferenceName string = "Go Conference"
var remainingTickets uint = conferenceTickets

// var bookings = make([]map[string]string, 0) //works fine but holds only one kind of datatype

var bookings = make([]UserData, 0)

//to store mixed data type in map like structure, structs are used

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {
	greetUsers()

	/*
		synchronizing goroutines. what if there was no for loop below,
		due to intoducing goroutines, the program would exit after taking
		the userdetails and would not wait for the ticket to be sent.

		thus - synchronizing go-routines.
	*/

	// for {
	userFirstName, userLastName, userEmail, userTickets := getUserInputs()

	isValidName, isValidEmail, isValidUserTicket := utils.ValidateUserInputs(userFirstName, userLastName, userEmail, userTickets, remainingTickets)

	if isValidName && isValidEmail && isValidUserTicket {
		bookTickets(userTickets, userFirstName, userLastName, userEmail)

		wg.Add(1) //the value inside the add function depends on the number of async funcs
		go sendTickets(userTickets, userFirstName, userLastName, userEmail)

		firstNames := getFirstNames()
		fmt.Printf("The first names of all of the bookings are %v\n", firstNames)

		noTicketsRemaining := remainingTickets == 0

		if noTicketsRemaining {
			fmt.Println("Our conference is booked out")
			// break
		}
	} else {
		fmt.Println("Your input data is invalid.")
	}

	wg.Wait()
	// }
}

func greetUsers() {
	fmt.Printf("Welcome to %v and the tickets left to buy is %v\n", conferenceName, conferenceTickets)
	fmt.Printf("We have %v tickets left\n", remainingTickets)
}

func getFirstNames() []string {
	firstNames := []string{}

	for _, booking := range bookings {
		// var namesArr []string = strings.Fields(booking)
		fn := booking.firstName
		firstNames = append(firstNames, fn)
	}

	return firstNames
}

func getUserInputs() (string, string, string, uint) {
	var userFirstName string
	var userLastName string
	var userEmail string
	var userTickets uint

	fmt.Println("Pease enter your first name")
	fmt.Scan(&userFirstName)
	fmt.Println("Please enter your last name")
	fmt.Scan(&userLastName)
	fmt.Println("Please enter your email address")
	fmt.Scan(&userEmail)
	fmt.Println("Please enter the number of tickets your want")
	fmt.Scan(&userTickets)

	return userFirstName, userLastName, userEmail, userTickets
}

func bookTickets(userTickets uint, userFirstName, userLastName, userEmail string) {
	remainingTickets = remainingTickets - userTickets

	//create a map for a user

	userData := UserData{
		firstName:       userFirstName,
		lastName:        userLastName,
		email:           userEmail,
		numberOfTickets: userTickets,
	}

	// userData["firstName"] = userFirstName
	// userData["lastName"] s= userLastName
	// userData["email"] = userEmail
	// userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	bookings = append(bookings, userData)

	fmt.Printf("Thank you %v %v for purchasing %v tickets. You will receive a confirmation mail at %v\n", userFirstName, userLastName, userTickets, userEmail)

	fmt.Printf("Tickets left - %v\n", remainingTickets)
	fmt.Printf("Bookings are %v\n", bookings)
}

func sendTickets(userTickets uint, firstName, lastName, email string) {
	time.Sleep(10 * time.Second)
	/**
	the above line blocks the main thread and hangs the program
	to handle situations like this, goroutines are used since
	concurrency is one of the biggest advantages of using go.
	*/
	ticket := fmt.Sprintf("%v tickets are booked for %v %v", userTickets, firstName, lastName)
	fmt.Println("-------------------")
	fmt.Printf("Sending ticket - %v to email - %v\n", ticket, email)
	fmt.Println("-------------------")

	wg.Done()
}
