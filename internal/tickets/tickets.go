package tickets

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
)

type Ticket struct {
	Id        string
	Nombre    string
	email     string
	Pais      string
	HoraVuelo string
	precio    string
}

func GetDataFromFile() ([]Ticket, error) {
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current directory:", err)
		return nil, err
	}

	filePath := filepath.Join(currentDir, "tickets.csv")

	file, err := os.Open(filePath)

	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV", err)
		return nil, err
	}

	var tickets []Ticket

	for _, record := range records {
		ticket := Ticket{
			record[0],
			record[1],
			record[2],
			record[3],
			record[4],
			record[5],
		}

		tickets = append(tickets, ticket)
	}
	return tickets, nil
}

// ejemplo 1
func GetTotalTickets(destination string, tickets []Ticket) (int, error) {
	var counter int
	for _, ticket := range tickets {
		if ticket.Pais == destination || destination == "Todos" {
			counter = counter + 1
		}
	}
	return counter, nil
}

// ejemplo 2
func GetMornings(time string) (int, error) {
	return 1, nil
}

// ejemplo 3
func AverageDestination(destination string, total int, data []Ticket) (float64, error) {
	ticketsByDest, _ := GetTotalTickets(destination, data)
	var averageDestination float64 = float64(ticketsByDest) * 100 / float64(total)
	return averageDestination, nil
}
