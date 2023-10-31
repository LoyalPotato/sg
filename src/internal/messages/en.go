package messages

const (
	Cleanup_Removing = "Removing config..."
	Cleanup_Success  = "Everything was reset."
	Cleanup_None     = "No config exists for that value."
	Cleanup_Error    = "Unable to run command. If error continues, try manual reset."

	NotGitEnv = "Not running in a git repository.\nUse git init to initialize a repository in the current folder."
	NoGitTown = "Git town is not installed.\nPlease follow the instructions on the installation page of git town."

	Setup_TokenInfo   = "In order to make changes to your repository we'll need to set a configuration with an access token.\nThis can be done by going in this link %s and following the instructions.\n"
	Setup_TokenInsert = "Insert Token: "
	Setup_Finish      = "Setup concluded! To get started run '%s'\n"
	Setup_Error       = "There was an issue performing the setup."

	Start_No_Config = "The correct config was not found.\nBe sure that you've run '%s' first to make sure everything is configured correctly.\n"
)
