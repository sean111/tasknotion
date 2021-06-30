package cmd

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use: "tasknotion",
		Short: "Notion backed task manager",
		Long: "A task manager based on taskwarrior that uses Notion for it's backend",
	}
)

func Execute() error {
	return rootCmd.Execute()
}