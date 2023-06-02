package main

import (
	"fmt"

	"main.go/internal/tickets"
)

// import (
// 	"github.com/bootcamp-go/desafio-go-bases"
// )

func main() {
	totalByCountry, err := tickets.GetTotalTickets("Indonesia")
	if err != nil {
		fmt.Println("Errorrrr")
	} else {
		fmt.Println(totalByCountry)
	}

	total, err := tickets.GetTotalTickets("Todos")
	value, err := tickets.AverageDestination("Brazil", total)

	if err != nil {
		fmt.Println("Errorrrr")
	} else {
		fmt.Println(value)
	}

}
