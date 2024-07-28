package techpalace

import (
	"fmt"
	"strings"
)

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return fmt.Sprintf("Welcome to the Tech Palace, %s",strings.ToUpper(customer))
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	s := strings.Repeat("*", numStarsPerLine)

	return s+"\n"+welcomeMsg+"\n"+s
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	
	t:=strings.ReplaceAll(oldMsg, "*","")
	return strings.TrimSpace(t)
}
