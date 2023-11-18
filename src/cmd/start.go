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

func startCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:               "start",
		Short:             startShort,
		Long:              long(startShort, startLong),
		DisableAutoGenTag: true,
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
	return cmd
}

type Start struct {
	mainBranch     string
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

	mainBranch, _ := gittown.GetMainBranchConfig()

	return Start{
		mainBranch:     utils.RemoveNonPrintables(mainBranch),
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
	start.feature()
	start.shipRefactor()
	start.finish()
}

func (s *Start) printStory() {
	fmt.Println(messages.Start_Story)
}

func (s *Start) hack(match string, branchName string) {
	fmt.Printf(messages.Start_Hack, styles.Faint(s.mainBranch), styles.Green(match))
	cmdMatch(match)
	gittown.Hack(branchName)
	fmt.Printf(messages.Start_Hack_Finished, styles.Faint(s.mainBranch))
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
	fmt.Printf(messages.Start_Bugfix_Tree, s.mainBranch)
	fmt.Println(messages.Start_Bugfix_Changes)
	confirmChanges()
	s.newPr()
	fmt.Printf(messages.Start_Bugfix_Finished, styles.Green(s.bugfixBranch), styles.Green(s.refactorBranch), styles.Green(s.refactorBranch))
}

func (s *Start) feature() {
	fmt.Printf(messages.Start_Feature, styles.Green(s.featureAppend))
	cmdMatch(s.featureAppend)
	gittown.Append(s.featureBranch)
	fmt.Printf(messages.Start_Feature_Tree, s.mainBranch)
	fmt.Printf(messages.Start_Feature_Refactor, styles.Green(s.refactorBranch))

	// Making changes to refactor
	var err error
	err = git.Checkout(s.refactorBranch)
	if err != nil {
		utils.Exit(fmt.Sprintf(messages.Generic_Error, err), 1)
	}
	fmt.Printf(messages.Start_Feature_Refactor_Changes, styles.Green(s.refactorBranch))
	confirmChanges()
	s.commitAll()
	err = git.Push()
	if err != nil {
		utils.Exit(fmt.Sprintf(messages.Generic_Error, err), 1)
	}
	// Changes finished

	err = git.Checkout(s.featureBranch)
	if err != nil {
		utils.Exit(fmt.Sprintf(messages.Generic_Error, err), 1)
	}

	fmt.Printf(messages.Start_Feature_Sync,
		styles.Green(s.featureBranch),
		styles.Green(messages.GitTown_Sync),
		styles.Green(s.bugfixBranch),
		styles.Green(s.featureBranch),
	)

	cmdMatch(messages.GitTown_Sync)

	err = gittown.Sync()
	if err != nil {
		utils.Exit(fmt.Sprintf(messages.Generic_Error, err), 1)
	}

	fmt.Println(messages.Start_Feature_Changes)
	confirmChanges()
	s.newPr()
	fmt.Println(messages.Start_Feature_Finished)
}

func (s *Start) shipRefactor() {
	fmt.Printf(messages.Start_Ship,
		styles.Faint(s.mainBranch),
		styles.Green("ship"),
		styles.Green(s.refactorShip),
	)
	cmdMatch(s.refactorShip)

	err := gittown.Ship(s.refactorBranch)

	if err != nil {
		utils.Exit(fmt.Sprintf(messages.Generic_Error, err), 1)
	}
}

func (s *Start) finish() {
	fmt.Printf(messages.Start_Finished,
		styles.FaintItalic("https://github.com/LoyalPotato/sg"),
		styles.Green("sync"))
}

func (s *Start) newPr() {
	s.commitAll()
	if err := gittown.PullRequest(); err != nil {
		utils.Exit(fmt.Sprintf(messages.Generic_Error, err), 1)
	}
}

func (s *Start) commitAll() {
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

func confirmChanges() {
	dialog.Confirm("Finished")
	changes, _ := git.GetNumOfChanges()
	for changes < 1 {
		fmt.Println(messages.No_Changes)
		dialog.Confirm("Finished")
		changes, _ = git.GetNumOfChanges()
	}
}
