/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"github.com/elishambadi/cli-todo-go/cmd"
)

func main() {
	// util.ReadCSVFile("todos.csv")
	// rowData := []string{"Play with the kids from Njeri's family", "2/2/2024", "no", "null"}
	// util.WriteOneRowToFile("todos.csv", rowData)
	// util.ReadCSVFile("todos.csv")
	cmd.Execute()
}
