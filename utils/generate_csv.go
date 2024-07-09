package utils

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Employee struct {
	ID  string
	Age int
}

func GenerateCSV() {
	records := []Employee{
		{"E01", 25},
		{"E02", 26},
		{"E03", 24},
		{"E04", 26},
	}
	csvFile, err := os.Create("employee.csv")
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	csvwriter := csv.NewWriter(csvFile)
	defer csvwriter.Flush()

	// Using WriteAll
	var data [][]string
	for _, record := range records {
		row := []string{record.ID, strconv.Itoa(record.Age)}
		data = append(data, row)
	}
	fmt.Println(data)
	csvwriter.WriteAll(data)

	csvFile.Close()
}
