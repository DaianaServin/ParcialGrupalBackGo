package main

import (
	"fmt"

	"main.go/internal/tickets"
)

// import (
// 	"github.com/bootcamp-go/desafio-go-bases"
// )

func main() {

	data, err := tickets.GetDataFromFile()

	if err != nil {
		fmt.Println("Error al traer la data")
		return
	}

	totalByCountry, err := tickets.GetTotalTickets("Indonesia", data)
	if err != nil {
		fmt.Println("Error en el calculo, intentelo de nuevo")
	} else {
		fmt.Println(totalByCountry)
	}

	total, err := tickets.GetTotalTickets("Todos", data)
	value, err := tickets.AverageDestination("Brazil", total, data)
	if err != nil {
		fmt.Println("Error en el calculo, intentelo de nuevo")
	} else {
		fmt.Println(value)
	}

}
