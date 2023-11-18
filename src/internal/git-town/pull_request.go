package gittown

// Calls hack cmd of git-town.
func PullRequest() error {
	return runGitTown("new-pull-request")
}
