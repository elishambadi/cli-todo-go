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

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete [id]",
	Short: "Delete a ToDo",
	Long:  `Deletes a task based on the ID provided.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("delete called")

		rowToDeleteId, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Error converting id to string")
			return
		}

		deleteError := util.DeleteRow(constants.STORAGE_FILE, rowToDeleteId)
		if deleteError != nil {
			fmt.Println(deleteError)
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
