package cmd

import (
	"fmt"
	jsondbprovider "todo/JsonDbProvider"
	model "todo/Model"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new task",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var taskName string = args[0]
		task := model.Task{Content: taskName, Complete: false}
		jsondbprovider.AddTask(model.Task(task))
		fmt.Println("Added new task " + taskName)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
