package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

func init() {
	rootCmd.AddCommand(newCommand)
}

var newCommand = &cobra.Command{
	Use: "new",
	Short: "Add a task",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Add a new task")
		fmt.Println(os.Getenv("DATABASE_ID"))
	},
}