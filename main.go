package main

import (
	"fmt"

	"github.com/gobwas/glob/util/strings"
	"main.go/internal/tickets"

	"time"
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

	go func(p string) {
		totalByCountry, err := tickets.GetTotalTickets("Indonesia", data)
	if err != nil {
		fmt.Println("Error en el calculo, intentelo de nuevo")
	} else {
		fmt.Println(totalByCountry)
	}
	}("Indonesia")

	totalByCountry, err := tickets.GetTotalTickets("Indonesia", data)
	if err != nil {
		fmt.Println("Error en el calculo, intentelo de nuevo")
	} else {
		fmt.Println(totalByCountry)
	}

	go func(p string) {
		total, err := tickets.GetTotalTickets("Todos", data)
		
	}("Todos")

	go func (p string)  {
		valueHour, err := tickets.GetMornings("6:00", data)
		if err != nil {
			fmt.Println("Error, intentelo de nuevo")
		} else {
			fmt.Println(valueHour)
		}
		fmt.Printf("El horario del vuelo es por la %d",data)
	} ("6:00")
	

	go func(p string) {
		value, err := tickets.AverageDestination("Brazil", total, data)
	if err != nil {
		fmt.Println("Error en el calculo, intentelo de nuevo")
	} else {
		fmt.Println(value)
	}
	}("Brazil", total)



}
