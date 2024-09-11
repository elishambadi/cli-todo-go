package file_utils

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"text/tabwriter"
	"time"

	tablewriter "github.com/olekukonko/tablewriter"
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

	if len(records) == 1 {
		fmt.Println("Woohoo! You have no tasks\n\nRun: ./cli-todo-go create [task_description] [end-date {DD-MM-YYYY}].\n")
		return
	}

	// Instantializes a new tabwriter
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(records[0])

	writer := tabwriter.NewWriter(os.Stdout, 0, 8, 2, ' ', 0)
	defer writer.Flush()

	for _, record := range records[1:] {
		// Check only for completed tasks.
		if completedOnly {
			if record[4] == "❌" {
				continue
			}
		}

		table.Append(record)
	}

	table.Render() // Displays the table
}

func WriteOneRowToFile(fileName string, rowData []string) {
	// Format 2nd param to date object
	dateFormat := "02-01-2006"
	inputDate := rowData[1]

	parsedDate, err1 := time.Parse(dateFormat, inputDate)
	if err1 != nil {
		fmt.Println("Please input the right date format {DD-MM-YYYY}.")
		return
	}

	rowData[1] = parsedDate.Format("Jan 02, 2006")

	// Append createdAt & time validation
	createdAt := time.Now().Format("Jan 02, 2006")
	if parsedDate.Unix() < time.Now().Unix() {
		fmt.Printf("\nYou cannot create a task in the past on %s. Today is %s\n\n", rowData[1], createdAt)
		return
	}
	rowData = append(rowData, createdAt)

	// ----- File opening ----
	err := os.Chmod(fileName, 0666)
	if err != nil {
		log.Fatalf("Failed to change file permissions: %s", err)
	}

	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		fmt.Print("Error opening file: ", err)
		return
	}

	// ---- Read file to get last ID ----
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
	if len(finalRowData) == 4 {
		paddingRecords := []string{"❌", "null"}
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

func MarkAsComplete(fileName string, rowId int) error {
	file, err := os.OpenFile(fileName, os.O_RDWR, 0666)
	if err != nil {
		return fmt.Errorf("error opening file: %s", err)
	}

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		return fmt.Errorf("error reading from file: %s", err)
	}

	updatedRecords := [][]string{
		{"id", "task", "created at", "due date", "completed", "user_id"},
	}

	recordFound := false
	for _, record := range records[1:] {
		if len(record) == 0 {
			continue
		}

		recordInt, err := strconv.Atoi(record[0])
		if err != nil {
			return fmt.Errorf("error converting recordId to string: %s", err)
		}

		if recordInt == rowId {
			record[4] = "✔️" // mark as complete
			recordFound = true
		}

		updatedRecords = append(updatedRecords, record)
	}

	if !recordFound {
		ReadCSVFile(fileName, false)
		return fmt.Errorf("\n-no record found with the given id %d", rowId)
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

	fmt.Println("Task completed Successfully successfully! +1 Karma. ")

	ReadCSVFile(fileName, false)

	return nil
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

	updatedRecords := [][]string{
		{"id", "task", "created_at", "due date", "completed", "user_id"},
	}

	recordFound := false
	for _, record := range records[1:] {
		if len(record) == 0 {
			continue
		}

		recordInt, err := strconv.Atoi(record[0])
		if err != nil {
			return fmt.Errorf("error converting recordId to string: %s", err)
		}

		if recordInt == rowId {
			recordFound = true
			continue // moves to next iter
		}

		updatedRecords = append(updatedRecords, record)
	}

	if !recordFound {
		ReadCSVFile(fileName, false)
		return fmt.Errorf("\n-no record found with the given id %d", rowId)
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

	fmt.Println("Task deleted successfully.")

	ReadCSVFile(fileName, false)

	return nil
}
