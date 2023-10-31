package styles

import "github.com/manifoldco/promptui"

var (
	Bold        = promptui.Styler(promptui.FGBold)
	Italic      = promptui.Styler(promptui.FGItalic)
	Faint       = promptui.Styler(promptui.FGFaint)
	FaintItalic = promptui.Styler(promptui.FGFaint, promptui.FGItalic)
)
