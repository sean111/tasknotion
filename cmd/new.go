package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"tasknotion/backend"
)

func init() {
	rootCmd.AddCommand(newCommand)
}

var newCommand = &cobra.Command{
	Use: "new",
	Short: "Add a task",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Add a new task")
		backend.Test()
	},
}