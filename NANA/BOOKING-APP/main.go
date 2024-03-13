package main

import "fmt"

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets = 50

	//fmt.Println("Welcome to our", conferenceName, "booking application") //with whitespaces for Println
	fmt.Printf("Welcome to our %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available!\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets to attend!")

}
