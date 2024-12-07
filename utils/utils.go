package utils

import "strings"

func ValidateUserInputs(userFirstName, userLastName, email string, userTickets, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(userFirstName) >= 2 && len(userLastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidUserTicket := userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidEmail, isValidUserTicket
}
