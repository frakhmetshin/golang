package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Conference" //var conferenceName string = "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	bookings := []string{}

	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)

	//fmt.Println("Welcome to our", conferenceName, "booking application") //with whitespaces for Println
	fmt.Printf("Welcome to our %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available!\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets to attend!")

	//var bookings = [50]string{}
	//var bookings [50]string //array type
	//var bookings []string // slice type
	for {
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

		remainingTickets = remainingTickets - userTickets
		//bookings[0] = firstName + " " + lastName //for array
		bookings = append(bookings, firstName+" "+lastName)

		//fmt.Printf("The whole slice: %v \n", bookings)
		//fmt.Printf("The first value: %v\n", bookings[0])
		//fmt.Printf("The type of slice:%T\n", bookings)
		//fmt.Printf("Array length: %v\n", len(bookings))
		fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
		fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
		firstNames := []string{}
		for _, booking := range bookings {
			var names = strings.Fields(booking)
			//var firstName = names[0]
			firstNames = append(firstNames, names[0])
		}
		//fmt.Printf("These are all our bookings: %v\n", bookings)
		fmt.Printf("These are all first names of our bookings: %v\n", firstNames)

	}

}
