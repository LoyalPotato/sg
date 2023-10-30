package cmd

import (
	"github.com/spf13/cobra"
)

func Execute() error {
	rootCmd := rootCmd()
	rootCmd.AddCommand(cleanup)

	rootCmd.AddGroup(&cobra.Group{
		ID:    "setup",
		Title: "Setup:",
	}, &cobra.Group{
		ID:    "cleanup",
		Title: "Cleanup: ",
	})

	return rootCmd.Execute()
}

func long(summary string, desc ...string) string {
	if len(desc) == 1 {
		return summary + ".\n" + desc[0]
	}
	return summary + "."
}
