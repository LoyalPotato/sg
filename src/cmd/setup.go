package cmd

import (
	"errors"
	"fmt"
	"os"

	gittown "github.com/LoyalPotato/sg/src/internal/git-town"
	"github.com/LoyalPotato/sg/src/internal/messages"
	"github.com/LoyalPotato/sg/src/internal/styles"
	"github.com/LoyalPotato/sg/src/internal/utils"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

const setupShort = "Setup everything to start the guide."

const setupLong = `
This includes:

- Github access token (use: git town)
- What the main branch is (use: git town)
`

func setupCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:               "setup",
		Short:             setupShort,
		Long:              long(setupShort, setupLong),
		Run:               runSetup,
		DisableAutoGenTag: true,
		GroupID:           "management",
	}

	return cmd
}

func runSetup(_ *cobra.Command, _ []string) {
	fmt.Printf(messages.Setup_TokenInfo, styles.FaintItalic("https://github.com/settings/tokens"))
	prompt := getTokenPrompt()
	token, err := prompt.Run()
	if err != nil {
		fmt.Println(messages.Setup_Error)
		os.Exit(1)
	}
	branchName := getMainBranch()
	err = gittown.AddGithubToken(token, true)
	if err != nil {
		utils.Exit(fmt.Sprintf(messages.Generic_Error, err), 1)
	}

	err = gittown.AddMainBranchConfig(branchName, true)
	if err != nil {
		utils.Exit(fmt.Sprintf(messages.Generic_Error, err), 1)
	}
	fmt.Printf("\n"+messages.Setup_Finish, styles.Green("sg start"))
}

func getTokenPrompt() promptui.Prompt {
	return promptui.Prompt{
		Label: messages.Setup_TokenInsert,
		Mask:  ' ',
		Templates: &promptui.PromptTemplates{
			Success: " ",
		},
		Validate: func(s string) error {
			if s == "" {
				return errors.New(messages.Empty_String)
			}

			return nil
		},
	}
}

func getMainBranch() string {
	fmt.Println(messages.Setup_MainBranchInfo)
	prompt := promptui.Prompt{
		Label:       messages.Setup_MainBranchConfig,
		Default:     "main",
		AllowEdit:   true,
		HideEntered: true,
		Validate: func(s string) error {
			if s == "" {
				return errors.New(messages.Empty_String)
			}

			return nil
		},
	}

	branch, err := prompt.Run()
	if err != nil {
		fmt.Println(messages.Setup_Error)
		os.Exit(1)
	}

	return branch
}
