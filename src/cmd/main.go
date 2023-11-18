package cmd

import (
	"fmt"

	"github.com/LoyalPotato/sg/src/internal/messages"
	"github.com/LoyalPotato/sg/src/internal/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

func Execute() error {
	rootCmd := rootCmd()
	rootCmd.AddCommand(setupCmd())
	rootCmd.AddCommand(cleanupCmd())
	rootCmd.AddCommand(startCmd())

	rootCmd.AddGroup(&cobra.Group{
		ID:    "management",
		Title: "Management:",
	})

	header := &doc.GenManHeader{
		Title:   "SG",
		Section: "1",
	}
	if err := doc.GenManTree(&rootCmd, header, "/tmp"); err != nil {
		utils.Exit(fmt.Sprintf(messages.Generic_Error, err), 1)
	}
	if err := doc.GenMarkdownTree(&rootCmd, "/tmp"); err != nil {
		utils.Exit(fmt.Sprintf(messages.Generic_Error, err), 1)
	}

	return rootCmd.Execute()
}

func long(summary string, desc ...string) string {
	if len(desc) == 1 {
		return summary + ".\n" + desc[0]
	}
	return summary + "."
}
