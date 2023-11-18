package git

import (
	"fmt"
	"strings"

	"github.com/LoyalPotato/sg/src/internal/cli"
	"github.com/LoyalPotato/sg/src/internal/messages"
	"github.com/LoyalPotato/sg/src/internal/utils"
)

func GetNumOfChanges() (int, error) {
	output, err := cli.RunCmdWithOutput("git", "status", "-s", "--porcelain")
	if err != nil {
		return 0, err
	}

	changes, err := utils.LineCounter(strings.NewReader(string(output)))
	if err != nil {
		return 0, err
	}

	return changes, nil
}

func AddAllChanges() error {
	fmt.Println(messages.Git_Add_All)
	return cli.RunCmd("git", "add", "-A")
}
