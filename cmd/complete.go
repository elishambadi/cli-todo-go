/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/elishambadi/cli-todo-go/constants"
	util "github.com/elishambadi/cli-todo-go/file_utils"
	"github.com/spf13/cobra"
)

// completeCmd represents the complete command
var completeCmd = &cobra.Command{
	Use:   "complete [id]",
	Short: "Mark a ToDo as complete",
	Args:  cobra.ExactArgs(1),
	Long: `Mark a task as complete.
	
	Provide the id as a param to declare that you have completed the task.`,
	Run: func(cmd *cobra.Command, args []string) {
		taskToCompleteId, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Error converting id to string")
			return
		}

		completeError := util.MarkAsComplete(constants.STORAGE_FILE, taskToCompleteId)
		if completeError != nil {
			fmt.Println(completeError)
		}
	},
}

func init() {
	rootCmd.AddCommand(completeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// completeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// completeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
