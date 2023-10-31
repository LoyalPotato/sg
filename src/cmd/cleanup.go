package cmd

import (
	"fmt"
	"os"

	gittown "github.com/LoyalPotato/sg/src/internal/git-town"
	"github.com/LoyalPotato/sg/src/internal/messages"
	"github.com/spf13/cobra"
)

const cleanupDesc = "Use this to reset what setup did."

const cleanupInfo = `
Here's a list of what that is:

- Remove token from git config
`

var cleanup = &cobra.Command{
	Use:     "cleanup",
	Short:   cleanupDesc,
	Long:    long(cleanupDesc, cleanupInfo),
	GroupID: "cleanup",
	Run: func(cmd *cobra.Command, args []string) {
		runCleanup()
	},
}

func runCleanup() {
	fmt.Println(messages.Cleanup_Removing)
	exists, err := gittown.RemoveGithubToken()
	if err != nil {
		fmt.Println(messages.Cleanup_Error)
		os.Exit(1)
	}

	msg := messages.Cleanup_Success

	if !exists {
		msg = messages.Cleanup_None
	}

	fmt.Println(msg)
}
