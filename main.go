package main

import (
	"fmt"
	"time"

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

	go func(p string) {
		totalByCountry, err := tickets.GetTotalTickets(p, data)
	if err != nil {
		fmt.Println("Error en el calculo, intentelo de nuevo")
	} else {
		fmt.Printf("El total para el pais %s es %d \n", p, totalByCountry)
	}
	}("Indonesia")
	time.Sleep(1 *time.Second)
	/* totalByCountry, err := tickets.GetTotalTickets("Indonesia", data)
	if err != nil {
		fmt.Println("Error en el calculo, intentelo de nuevo")
	} else {
		fmt.Println(totalByCountry)
	} */

	go func(p string) {
		total, err := tickets.GetTotalTickets(p, data)
		if err != nil{
			fmt.Println(err)
		}
		fmt.Printf("El total es %d \n", total)
		
	}("Todos")
	time.Sleep(1 *time.Second)

	go func (p string)  {
		valueHour, err := tickets.GetMornings(p, data)
		if err != nil {
			fmt.Println("Error, intentelo de nuevo")
		} 
		fmt.Println("La cantidad de pasajeros es de : ", valueHour)
	} ("6:00")
	time.Sleep(1 *time.Second)

	go func(p string) {
		total, err := tickets.GetTotalTickets("Todos", data)
		if err != nil{
			fmt.Println(err)
		}

		value, err := tickets.AverageDestination(p, total, data)
		if err != nil {
			fmt.Println("Error en el calculo, intentelo de nuevo")
		} else {
			fmt.Printf("El porcentaje para el %s es %.2f ", p, value)
		}
	}("Brazil")
	time.Sleep(1 *time.Second)

}
