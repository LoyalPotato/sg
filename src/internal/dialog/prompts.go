package dialog

import (
	"errors"
	"os"

	"github.com/LoyalPotato/sg/src/internal/messages"
	"github.com/LoyalPotato/sg/src/internal/utils"
	"github.com/manifoldco/promptui"
)

type MatchPromptArgs struct {
	Label    string
	MatchVal string
}

// Creates a prompt that is valid when the value is the same as the default.
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
// Returns bool value, if confirmed, and error value.
func Confirm(label string) (bool, error) {
	prompt := promptui.Prompt{
		Label:       label,
		IsConfirm:   true,
		HideEntered: true,
	}

	_, err := prompt.Run()

	if errors.Is(err, promptui.ErrAbort) {
		return false, err
	} else if errors.Is(err, promptui.ErrInterrupt) {
		os.Exit(utils.ExitCodeInterrupt)
	}

	return true, err
}

type TextPromptArgs struct {
	Label string

	Default string

	AllowEdit bool

	Validate promptui.ValidateFunc

	Mask rune

	HideEntered bool

	Templates *promptui.PromptTemplates
}

// Text prompt.
func Text(args *TextPromptArgs) (output string, err error) {
	prompt := promptui.Prompt{
		Label:       args.Label,
		AllowEdit:   args.AllowEdit,
		HideEntered: args.HideEntered,
		Mask:        args.Mask,
		Validate:    args.Validate,
		Templates:   args.Templates,
	}

	output, err = prompt.Run()

	return output, err
}
