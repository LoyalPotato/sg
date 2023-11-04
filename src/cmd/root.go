package cmd

import (
	"fmt"
	"os"

	"github.com/LoyalPotato/sg/src/internal/git"
	gittown "github.com/LoyalPotato/sg/src/internal/git-town"
	"github.com/LoyalPotato/sg/src/internal/messages"
	"github.com/spf13/cobra"
)

const rootDesc = "Interactive guide to learn stacking workflow"

const rootInfo = `
Hate big pull requests? Me too :)
Want to know how to make them smaller and not have to deal with all the git commands?

Then this tool will help you with just all that!

It's a step by step guide that will guide you through the git-town commands that are used to follow
the stacking workflow.

To learn more about the tool used, check the official git-town documentation: https://www.git-town.com/
And to learn more about the stacking workflow: https://stacking.dev`

func rootCmd() cobra.Command {
	rootCmd := cobra.Command{
		Use:               "sg",
		Short:             rootDesc,
		Long:              long(rootDesc, rootInfo),
		SilenceErrors:     true,
		SilenceUsage:      true,
		DisableAutoGenTag: true,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			if err := git.IsGitRepo(); err != nil {
				fmt.Println(messages.NotGitEnv)
				os.Exit(1)
			}

			if !gittown.IsInstalled() {
				fmt.Println(messages.NoGitTown)
				os.Exit(1)
			}
		},
	}

	return rootCmd
}
