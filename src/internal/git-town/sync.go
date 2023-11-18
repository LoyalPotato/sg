package gittown

// Calls append cmd of git-town.
func Sync() error {
	return runGitTown("sync")
}
