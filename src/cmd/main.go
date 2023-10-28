package cmd

import "github.com/spf13/cobra"

func Execute() error {
	rootCmd := rootCmd()

	rootCmd.AddGroup(&cobra.Group{
		ID:    "setup",
		Title: "Setup:",
	})

	return rootCmd.Execute()
}

func long(summary string, desc ...string) string {
	if len(desc) == 1 {
		return summary + ".\n" + desc[0]
	}
	return summary + "."
}
