package main

import (
	"encoding/csv"
	"os"
)

func writeLineToCSVFile(filename string, records [][]string) {
	file, err := os.Create(filename)
	errExit(err)
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, record := range records {
		err := writer.Write(record)
		errExit(err)
	}
}
