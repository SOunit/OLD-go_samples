package main

import "strings"

func ValidateUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := 0 < userTickets && userTickets <= remainingTickets

	return isValidName, isValidEmail, isValidTicketNumber
}
