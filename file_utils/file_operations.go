package file_utils

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"text/tabwriter"
)

// readCSVFile reads the contents of a file specified by filename and returns it as a string.
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
func ReadCSVFile(fileName string, completedOnly bool) {
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

	fmt.Println("------------------- Your Tasks -------------------- ")

	// Instantializes a new tabwriter
	writer := tabwriter.NewWriter(os.Stdout, 0, 8, 2, ' ', 0)
	defer writer.Flush()

	for _, record := range records {
		tabbedString := ""

		// Check only for completed tasks.
		if completedOnly {
			if record[3] == "no" {
				continue
			}
		}

		for _, column := range record {
			// fmt.Println(column)
			tabbedString += column + "\t"
		}
		fmt.Fprintln(writer, tabbedString)
	}
}

func WriteOneRowToFile(fileName string, rowData []string) {
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
		if err != nil {
			log.Fatalf("Error converting string to int in writeOneRowToFile function: %s", err)
		}

		nextRecordId = append(nextRecordId, strconv.Itoa(nextRecordIdInt+1))
	}
	// Append the ID
	finalRowData := Prepend(rowData, nextRecordId)

	// Padd the remaining tasks if the len is too small
	if len(finalRowData) == 3 {
		paddingRecords := []string{"no", "null"}
		finalRowData = append(finalRowData, paddingRecords...)
	}

	writer := csv.NewWriter(file)

	writer.Write(finalRowData)

	writer.Flush() // Write buffered data to file

	if err := writer.Error(); err != nil {
		log.Fatalf("error flushing data to file: %s", err)
	}

	fmt.Printf("Created a task \"%s\" successfully. Due on %s. \n\nCall 'list' command to view all tasks.\n", finalRowData[1], finalRowData[2])
	file.Close()
}

func DeleteRow(fileName string, rowId int) error {

	file, err := os.OpenFile(fileName, os.O_RDWR, 0666)
	if err != nil {
		return fmt.Errorf("error opening file: %s", err)
	}

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		return fmt.Errorf("error reading from file: %s", err)
	}

	updatedRecords := [][]string{}

	for _, record := range records[1:] {
		if len(record) == 0 {
			continue
		}

		recordInt, err := strconv.Atoi(record[0])
		if err != nil {
			return fmt.Errorf("error converting recordId to string: %s", err)
		}

		if recordInt == rowId {
			continue // moves to next iter
		}

		updatedRecords = append(updatedRecords, record)
	}
	file.Close()

	updatedFile, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("error creating a new file for writing: %s", err)
	}

	//Write the updated records to file
	writer := csv.NewWriter(updatedFile)

	if writeError := writer.WriteAll(updatedRecords); writeError != nil {
		return fmt.Errorf("error writing updated records to file: %s", writeError)
	}

	fmt.Println("Task ", records[rowId-1][1], " deleted successfully.")

	ReadCSVFile(fileName, false)

	return nil
}
