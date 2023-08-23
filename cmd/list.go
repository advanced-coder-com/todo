package cmd

import (
	"fmt"
	jsondbprovider "todo/JsonDbProvider"
	model "todo/Model"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Print all tasks in the list",
	Run: func(cmd *cobra.Command, args []string) {
		_PrintList()
	},
}

func _PrintList() {
	var tasks model.Tasks = *jsondbprovider.GetList()

	for key, val := range tasks.Tasks {
		if !val.Complete {
			fmt.Println(key, val.Content)
		}
	}
}

func init() {
	rootCmd.AddCommand(listCmd)
}
