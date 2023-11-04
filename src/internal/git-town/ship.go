package gittown

func Ship(branch string) error {
	return runGitTown("ship", branch)
}
