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
	Pais      string
	HoraVuelo string
	precio    string
}

func GetDataFromFile() []Ticket {
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current directory:", err)
		return nil
	}

	filePath := filepath.Join(currentDir, "tickets.csv")

	file, err := os.Open(filePath)

	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV", err)
		return nil
	}

	var tickets []Ticket

	for _, record := range records {
		ticket := Ticket{
			record[0],
			record[1],
			record[2],
			record[3],
			record[4],
		}

		tickets = append(tickets, ticket)
	}
	return tickets
}

// ejemplo 1
func GetTotalTickets(destination string) (int, error) {
	var tickets []Ticket
	tickets = GetDataFromFile()
	var counter int
	for i, _ := range tickets {
		counter = i + 1
	}
	return counter, nil
}

// ejemplo 2
func GetMornings(time string) (int, error) {
	return 1, nil
}

// ejemplo 3
func AverageDestination(destination string, total int) (int, error) {
	return 1, nil
}
