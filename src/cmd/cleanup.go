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

- Remove token from git config (use: git town)
- Remove main branch config (use: git town)
`

func cleanupCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:               "cleanup",
		Short:             cleanupDesc,
		Long:              long(cleanupDesc, cleanupInfo),
		GroupID:           "management",
		DisableAutoGenTag: true,
		Run: func(cmd *cobra.Command, args []string) {
			runCleanup()
		},
	}
	return cmd
}

func runCleanup() {
	fmt.Println(messages.Cleanup_Removing)
	exists, err := gittown.RemoveGithubToken()
	mainConfExists, mainConfErr := gittown.RemoveMainBranchConfig()
	if err != nil || mainConfErr != nil {
		fmt.Println(messages.Cleanup_Error)
		os.Exit(1)
	}

	msg := messages.Cleanup_Success

	if !exists && !mainConfExists {
		msg = messages.Cleanup_None
	}

	fmt.Println(msg)
}
