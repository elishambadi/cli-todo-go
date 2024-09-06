/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"encoding/csv"
	"fmt"
	"os"
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

	for i := 0; i < len(records); i++ {
		fmt.Println("Column: ", records[i])
	}
}

func writeToFile(fileName string, rowData []string) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Print("Error opening file: ", err)
	}

	writer := csv.NewWriter(file)

	writer.Write()
}

func main() {
	readCSVFile("todos.csv")
	// cmd.Execute()
}
