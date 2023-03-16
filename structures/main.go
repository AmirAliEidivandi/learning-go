package main

import "fmt"

type Sample struct {
	field string
	a, b  int
}

type Passenger struct {
	Name         string
	TicketNumber int
	Boarded      bool
}

type Bus struct {
	FrontSeat Passenger
}

func main() {
	casey := Passenger{"Casey", 1, false}
	fmt.Println(casey)

	var (
		bill = Passenger{Name: "Bill", TicketNumber: 5}
		ella = Passenger{Name: "Ella", TicketNumber: 5}
	)
	fmt.Println(bill, ella)

	var heidi Passenger
	heidi.Name = "heidi"
	heidi.TicketNumber = 120
	fmt.Println(heidi)

	casey.Boarded = true
	heidi.Boarded = true

	if bill.Boarded {
		fmt.Println("Bill has boarded the bus")
	}
	if casey.Boarded {
		fmt.Println(casey.Name, "has boarded the bus")
	}

	heidi.Boarded = true
	bus := Bus{heidi}
	fmt.Println(bus)
	fmt.Println(bus.FrontSeat.Name, "is in the front seat")

	// data := Sample{"world", 5, 4} // way 1
	data := Sample{
		field: "world",
		a:     5,
		b:     6,
	}
	fmt.Println(data.field, data.a, data.b)

	// accessing structures
	data.field = "bye bye"
	data.a = 20
	data.b = 50
	fmt.Println(data.field, data.a, data.b)

	// anonymous structures
	example := struct {
		field string
		a, b  int
	}{
		"hello",
		5, 3,
	}
	fmt.Println(example.a, example.b, example.field)
}
