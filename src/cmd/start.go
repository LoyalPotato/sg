package cmd

import (
	"fmt"
	"os"

	"github.com/LoyalPotato/sg/src/internal/dialog"
	"github.com/LoyalPotato/sg/src/internal/git"
	gittown "github.com/LoyalPotato/sg/src/internal/git-town"
	"github.com/LoyalPotato/sg/src/internal/messages"
	"github.com/LoyalPotato/sg/src/internal/styles"
	"github.com/LoyalPotato/sg/src/internal/utils"
	"github.com/spf13/cobra"
)

const startShort = "Command to start guide"

const startLong = `
You'll be guided through a story and learn how to apply the stacking workflow.

The story will teach you:
- Start feature
- Build on on-going features
- Sync all feature branches
- Ship feature branches
`

var start = &cobra.Command{
	Use:   "start",
	Short: startShort,
	Long:  long(startShort, startLong),
	PreRun: func(cmd *cobra.Command, args []string) {
		exists := git.ConfigExists(gittown.TokenKey)
		existsMain := git.ConfigExists(gittown.MainBranchKey)
		if !exists || !existsMain {
			fmt.Printf(messages.Start_No_Config, styles.Green("sg setup"))
			os.Exit(1)
		}
	},
	Run: runStart,
}

type Start struct {
	refactorBranch string
	bugfixBranch   string
	featureBranch  string

	refactorHack  string
	bugAppend     string
	featureAppend string
	refactorShip  string
}

func startCtor() Start {
	const refactorBranch = "1-refactor"
	const bugfixBranch = "2-bugfix"
	const featureBranch = "3-feature"

	return Start{
		refactorBranch: refactorBranch,
		bugfixBranch:   bugfixBranch,
		featureBranch:  featureBranch,

		refactorHack:  fmt.Sprintf(messages.GitTown_Hack, refactorBranch),
		refactorShip:  fmt.Sprintf(messages.GitTown_Ship, refactorBranch),
		bugAppend:     fmt.Sprintf(messages.GitTown_Append, bugfixBranch),
		featureAppend: fmt.Sprintf(messages.GitTown_Append, featureBranch),
	}
}

func runStart(cmd *cobra.Command, args []string) {
	start := startCtor()
	start.printStory()
	start.refactor()
	start.bugfix()

func (s *Start) printStory() {
	fmt.Println(messages.Start_Story)
}

func (s *Start) hack(match string, branchName string) {
	fmt.Printf(messages.Start_Hack, styles.Green(match))
	cmdMatch(match)
	gittown.Hack(branchName)
	fmt.Println(messages.Start_Hack_Finished)
}

func (s *Start) refactor() {
	s.hack(s.refactorHack, s.refactorBranch)
	fmt.Println(messages.Start_Refactor_Changes)
	confirmChanges()
	s.newPr()
	fmt.Println(messages.Start_Refactor_Finished)
}

func (s *Start) bugfix() {
	fmt.Printf(messages.Start_Bugfix, styles.Green(s.bugAppend))
	cmdMatch(s.bugAppend)
	gittown.Append(s.bugfixBranch)
	fmt.Println(messages.Start_Bugfix_Tree)
	fmt.Println(messages.Start_Bugfix_Changes)
	confirmChanges()
	s.newPr()
	fmt.Printf(messages.Start_Bugfix_Finished, styles.Green(s.bugfixBranch), styles.Green(s.refactorBranch), styles.Green(s.refactorBranch))
}

func (s *Start) newPr() {
	var err error

	err = git.AddAllChanges()

	if err != nil {
		utils.Exit(fmt.Sprintf(messages.Generic_Error, err), 1)
	}

	m, err := dialog.Text(&dialog.TextPromptArgs{
		Label:     messages.Git_Commit,
		Default:   messages.Git_Default_Commit_Msg,
		AllowEdit: false,
	})

	if err != nil {
		utils.Exit(fmt.Sprintf(messages.Generic_Error, err), 1)
	}

	if m == "" {
		m = messages.Git_Default_Commit_Msg
	}

	err = git.Commit(m)

	if err != nil {
		utils.Exit(fmt.Sprintf(messages.Error_Git_Commit, err), 1)
	}

	if gittown.PullRequest() != nil {
		utils.Exit(fmt.Sprintf(messages.Generic_Error, err), 1)
	}
}

func cmdMatch(match string) {
	_, err := dialog.Match(&dialog.MatchPromptArgs{
		Label:    "Command",
		MatchVal: utils.RemoveNonPrintables(match),
	})

	if err != nil {
		utils.Exit(fmt.Sprintf(messages.Generic_Error, err), 1)
	}
}

// TODO: This loop isn't working
// It won't exit with ^C
// After doing ^C and it saying no changes, if I make the changes finally, it's not reading correctly
func confirmChanges() {
	dialog.Confirm("Finished")
	for changes, _ := git.GetNumOfChanges(); changes < 1; {
		fmt.Println(messages.No_Changes)
		dialog.Confirm("Finished")
	}
}
