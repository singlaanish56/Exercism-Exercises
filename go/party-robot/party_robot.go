package partyrobot

import (
	"fmt"
)
// Welcome greets a person by name.
func Welcome(name string) string {
	x := fmt.Sprintf("Welcome to my party, %s!", name)
	return x
}

// HappyBirthday wishes happy birthday to the birthday person and exclaims their age.
func HappyBirthday(name string, age int) string {
	
	result := fmt.Sprintf("Happy birthday %s! You are now %d years old!", name, age)
	return result
} 

// AssignTable assigns a table to each guest.
func AssignTable(name string, table int, neighbor, direction string, distance float64) string {
	tableNumberStr := fmt.Sprintf("%03d", table)

	distanceStr := fmt.Sprintf("%.1f", distance)
	result := fmt.Sprintf("Welcome to my party, %s!\n", name)
	result += fmt.Sprintf("You have been assigned to table %s. Your table is %s, exactly %s meters from here.\n", tableNumberStr, direction, distanceStr)
	result += fmt.Sprintf("You will be sitting next to %s.", neighbor)

	return result
}
