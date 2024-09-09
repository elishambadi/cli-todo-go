/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

	"text/tabwriter"
)

// readFile reads the contents of a file specified by filename and returns it as a string.
//
// Parameters:
//   - filename: The name of the file to read.
//
// Returns:
//   - The contents of the file as a string.
//   - An error if the file could not be read.
//
// Usage example:
//
//	content, err := readFile("example.txt")
//	if err != nil {
//	    log.Fatalf("Error reading file: %v", err)
//	}
//	fmt.Println(content)
func readCSVFile(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("error opening the csv file: ", err)
		return
	}
	defer file.Close()

	// fmt.Println(file)

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading from your CSV file: ", err)
		return
	}

	fmt.Println("----------------------- File: ", fileName, " ---------------------------------------- ")

	// Instantializes a new tabwriter
	writer := tabwriter.NewWriter(os.Stdout, 0, 8, 2, ' ', 0)
	defer writer.Flush()

	for i, record := range records {
		fmt.Println("----------- Record ", i+1, "----------------")
		for _, column := range record {
			fmt.Println(column)
		}

	}

	for i := 0; i < len(records); i++ {
		fmt.Fprintln(writer, records[i])
		// fmt.Println("Column: ", records[i])
	}
}

func writeOneRowToFile(fileName string, rowData []string) {
	err := os.Chmod(fileName, 0666)
	if err != nil {
		log.Fatalf("Failed to change file permissions: %s", err)
	}

	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		fmt.Print("Error opening file: ", err)
	}

	// Read file to get last ID
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading from your CSV file in write function: ", err)
		return
	}

	nextRecordId := []string{}

	if len(records) == 1 {
		// 1 means only headers row exists
		nextRecordId = append(nextRecordId, "1")
	} else {
		nextRecordIdInt, err := strconv.Atoi(records[len(records)-1][0])
		log.Println("Next ID: ", nextRecordIdInt, " Record: ", records[len(records)-1][0])
		if err != nil {
			log.Fatalf("Error converting string to int in writeOneRowToFile function: %s", err)
		}

		nextRecordId = append(nextRecordId, strconv.Itoa(nextRecordIdInt+1))
		fmt.Println("Next record ID:", nextRecordId, nextRecordIdInt+1)
	}
	finalRowData := Prepend(rowData, nextRecordId)

	writer := csv.NewWriter(file)

	writer.Write(finalRowData)

	writer.Flush() // Write buffered data to file

	if err := writer.Error(); err != nil {
		log.Fatalf("error flushing data to file: %s", err)
	}

	log.Printf("written record successfully: %v,", rowData)
	file.Close()
}

func main() {
	readCSVFile("todos.csv")
	rowData := []string{"Change my engine oil", "2/2/2024", "no", "null"}
	writeOneRowToFile("todos.csv", rowData)
	readCSVFile("todos.csv")
	// cmd.Execute()
}
