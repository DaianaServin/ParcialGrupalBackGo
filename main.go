package main

import (
	"fmt"

	"main.go/internal/tickets"
)

// import (
// 	"github.com/bootcamp-go/desafio-go-bases"
// )

func main() {
	// total, err := tickets.GetTotalTickets("Brazil")
	value, err := tickets.GetTotalTickets("Indonesia")
	if err != nil {
		fmt.Println("Errorrrr")
	} else {
		fmt.Println(value)
	}

}
