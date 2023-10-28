package cmd

import "github.com/spf13/cobra"

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
		Use:           "sg",
		Aliases:       []string{"stacked-guide"},
		Short:         rootDesc,
		Long:          long(rootDesc, rootInfo),
		SilenceErrors: true,
		SilenceUsage:  true,
	}

	return rootCmd
}
