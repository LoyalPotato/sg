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

// Confirm prompt
//
// Returns bool value, if confirmed, and error value
func Confirm(label string) (bool, error) {
	prompt := promptui.Prompt{
		Label:       label,
		IsConfirm:   true,
		HideEntered: true,
	}
	_, err := prompt.Run()

	switch err {
	case promptui.ErrAbort:
		return false, err
	case promptui.ErrInterrupt:
		os.Exit(utils.EXIT_CODE_INTERRUPT)
	}
	return true, err
}

