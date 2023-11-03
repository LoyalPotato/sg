package cmd

import (
	"errors"
	"fmt"
	"os"

	gittown "github.com/LoyalPotato/sg/src/internal/git-town"
	"github.com/LoyalPotato/sg/src/internal/messages"
	"github.com/LoyalPotato/sg/src/internal/styles"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

const setupShort = "Setup everything to start the guide."

const setupLong = `
This includes:

- Github access token (use: git town)
- What the main branch is (use: git town)
`

var setup = &cobra.Command{
	Use:     "setup",
	Short:   setupShort,
	Long:    long(setupShort, setupLong),
	Run:     runSetup,
	GroupID: "setup",
}

func runSetup(cmd *cobra.Command, args []string) {
	fmt.Printf(messages.Setup_TokenInfo, styles.FaintItalic("https://github.com/settings/tokens"))
	prompt := getTokenPrompt()
	token, err := prompt.Run()
	if err != nil {
		fmt.Println(messages.Setup_Error)
		os.Exit(1)
	}
	branchName := getMainBranch()
	gittown.AddGithubToken(token, true)
	gittown.AddMainBranchConfig(branchName, true)
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
