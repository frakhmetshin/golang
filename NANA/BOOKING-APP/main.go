package main

import (
	"booking-app/helper"
	"fmt"
)

const conferenceTickets int = 50

var conferenceName = "Go Conference" //var conferenceName string = "Go Conference"
var remainingTickets uint = 50

// var bookings = make([]map[string]string, 0)
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

func main() {

	greetUsers()
	//fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)

	//fmt.Println("Welcome to our", conferenceName, "booking application") //with whitespaces for Println
	//fmt.Printf("Welcome to our %v booking application\n", conferenceName)
	//var bookings = [50]string{}
	//var bookings [50]string //array type
	//var bookings []string // slice type
	for {
		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			bookTicket(userTickets, firstName, lastName, email)
			fmt.Printf("These are all first names of our bookings: %v\n", getFirstNames())
			fmt.Printf("list of booking is %v \n", bookings)
			//noTicketsRemaining := remainingTickets == 0
			if remainingTickets == 0 {
				//end program
				fmt.Println("Our conbference is booked out! Come back next year")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("First name or last name is too short!")
			}
			if !isValidEmail {
				fmt.Println("The email address you entered doesn't contain @ sign!")
			}
			if !isValidTicketNumber {
				fmt.Println("Number of tickets you entered is invalid")
			}

			//continue
			//			city := "London"
			//			switch city {
			//			case "New York":
			//			//execute code
			//			case "Singapore","Hong Kong":
			//			//execute
			//			case "London","Berlin":
			//execute
			//			case "Mexico City":
			//execute
			//			default:
			//				fmt.Print("No Valid city selected")
			//			}

		}

	}

}

func greetUsers() {
	fmt.Printf("Welcome to our %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available!\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets to attend!")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		//var names = strings.Fields(booking)
		//var firstName = names[0]
		firstNames = append(firstNames, booking.firstName)
	}
	//fmt.Printf("These are all our bookings: %v\n", bookings)
	//fmt.Printf("These are all first names of our bookings: %v\n", firstNames)
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint
	//ask user for their name
	//fmt.Println(remainingTickets) see value of var
	//fmt.Println(&remainingTickets) see pointer of var
	//	userName = "Tom"

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
	//bookings[0] = firstName + " " + lastName //for array

	// create map for a user
	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}
	//	userData["firstName"] = firstName
	//	userData["lastName"] = lastName
	//	userData["email"] = email
	//	userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	bookings = append(bookings, userData)

	//fmt.Printf("The whole slice: %v \n", bookings)
	//fmt.Printf("The first value: %v\n", bookings[0])
	//fmt.Printf("The type of slice:%T\n", bookings)
	//fmt.Printf("Array length: %v\n", len(bookings))
	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}
