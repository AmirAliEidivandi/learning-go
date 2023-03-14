package main

import "fmt"

// Days of the week
const (
	Monday    = 0
	Tuesday   = 1
	Wednesday = 2
	Thursday  = 3
	Friday    = 4
	Saturday  = 5
	Sunday    = 6
)

// User roles
const (
	Admin      = 10
	Manager    = 20
	Contractor = 30
	Member     = 40
	Guest      = 50
)

func accessGranted() {
	fmt.Println("Granted")
}

func accessDenied() {
	fmt.Println("Denied")
}

func weekDay(day int) bool {
	return day <= 4
}

func main() {
	// today, role := Tuesday, Guest // Denied
	today, role := Sunday, Contractor // Granted

	if role == Admin || role == Manager {
		accessGranted()
	} else if role == Contractor && !weekDay(today) {
		accessGranted()
	} else if role == Member && weekDay(today) {
		accessGranted()
	} else if role == Guest && (today == Monday || today == Wednesday || today == Friday) {
		accessGranted()
	} else {
		accessDenied()
	}
}