package dialog

import (
	"errors"

	"github.com/LoyalPotato/sg/src/internal/messages"
	"github.com/LoyalPotato/sg/src/internal/utils"
	"github.com/manifoldco/promptui"
)

type MatchPromptArgs struct {
	Label    string
	MatchVal string
}

// Creates a prompt that is valid when the value is the same as the default
func Match(args *MatchPromptArgs) (string, error) {
	prompt := promptui.Prompt{
		Label: args.Label,
		Validate: func(s string) error {
			if s == "" {
				return errors.New(messages.Empty_String)
			}

			if !utils.StringMatches(s, args.MatchVal) {
				return errors.New("value doesn't match")
			}

			return nil
		},
		Templates: &promptui.PromptTemplates{},
	}

	res, err := prompt.Run()

	return res, err
}

