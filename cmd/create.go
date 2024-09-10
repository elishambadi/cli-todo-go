/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/elishambadi/cli-todo-go/constants"
	util "github.com/elishambadi/cli-todo-go/file_utils"
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create [task_name] [due_date {DD-MM-YYYY}]",
	Short: "Creates a new ToDo item",
	Long: `Command to create a new ToDo item.
	
	Expects 2 parameters: Task Name and the Due Date.`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		taskName := args[0]
		dueDate := args[1]

		taskData := []string{taskName, dueDate}

		util.WriteOneRowToFile(constants.STORAGE_FILE, taskData)
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
