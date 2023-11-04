package cmd

import (
	"github.com/spf13/cobra"
)

func Execute() error {
	rootCmd := rootCmd()
	rootCmd.AddCommand(setup)
	rootCmd.AddCommand(cleanup)
	rootCmd.AddCommand(start)

	rootCmd.AddGroup(&cobra.Group{
		ID:    "management",
		Title: "Management:",
	})

	return rootCmd.Execute()
}

func long(summary string, desc ...string) string {
	if len(desc) == 1 {
		return summary + ".\n" + desc[0]
	}
	return summary + "."
}
