/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"log"
	"strconv"
	jsondbprovider "todo/JsonDbProvider"

	"github.com/spf13/cobra"
)

// completeCmd represents the complete command
var completeCmd = &cobra.Command{
	Use:   "complete",
	Short: "Complete the task by ID",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var taskId string = args[0]
		intVar, err := strconv.Atoi(taskId)
		jsondbprovider.SetAsComplete(intVar)

		if err != nil {
			log.Fatal("Cannot parse task number:", err)
		}
		fmt.Println("Task " + taskId + " was completed")
	},
}

func init() {
	rootCmd.AddCommand(completeCmd)
}
