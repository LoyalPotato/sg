package gittown

// Calls append cmd of git-town.
func Append(branch string) error {
	return runGitTown("append", branch)
}
