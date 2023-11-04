package gittown

// Calls hack cmd of git-town
func Hack(branch string) error {
	return runGitTown("hack", branch)
}
