package messages

const (
	NotGitEnv              = "Not running in a git repository.\nUse git init to initialize a repository in the current folder."
	NoGitTown              = "Git town is not installed.\nPlease follow the instructions on the installation page of git town."
	Git_Commit             = "Insert commit message"
	Git_Default_Commit_Msg = "Default commit message"
	Git_Add_All            = "Staging all changes to current branch"

	GitTown_Hack   = "git-town hack %s\n"
	GitTown_Append = "git-town append %s\n"
	GitTown_Ship   = "git-town ship %s\n"

	No_Changes = "It seems like there are no new changes... Please make a change to be able to proceed!"

	Cleanup_Removing = "Removing config..."
	Cleanup_Success  = "Everything was reset."
	Cleanup_None     = "No config exists for that value."
	Cleanup_Error    = "Unable to run command. If error continues, try manual reset."

	Setup_TokenInfo        = "In order to make changes to your repository we'll need to set a configuration with an access token.\nThis can be done by going in this link %s and following the instructions.\n"
	Setup_TokenInsert      = "Insert Token"
	Setup_MainBranchInfo   = "Git town needs to know what is your main branch to then merge into and create from."
	Setup_MainBranchConfig = "What is your main branch"
	Setup_Finish           = "Setup concluded! To get started run '%s'\n"
	Setup_Error            = "There was an issue performing the setup."

	Start_No_Config = "The correct config was not found.\nBe sure that you've run '%s' first to make sure everything is configured correctly.\n"
	Start_Story     = `Let's start by defining the story for our workflow.
	
We have a new feature that we need to develop. But there are some other things that we need to do before starting the implementation:

- Clean up some code
- Fix a bug preventing the new feature from working 100%
- Implement the feature
	
We should clean up first, to have a good starting base!`
	Start_Hack = `Let's create a new branch from main to start working on the cleanup.
This can be done with the following command from git-town:
%s
Type it out below to run it :)
`
	Start_Hack_Finished = `Well done!
This command just created your new feature from the main branch you setup and checked it out.`
	Start_Refactor_Changes = `Now make some changes to any file.
Just something simple like creating a new file and writing "Refactoring", for example.
Once you're finished, proceed with the guide.`
	Start_Refactor_Finished = `Now all that's left for this step is to finish creating the PR, through the link opened.
If the link didn't open in your default browser, manually open it to conclude the creation of the PR.

And we've successfully put our first changes for review!`
	Start_Bugfix = `We can now start working on the bugfix on top of these changes while they are still in review.
To build on top of the previous changes, in the refactor branch still, we use the following git-town command:
%s
Give it a go:`
	Start_Bugfix_Tree = `This is what the branch hierarchy is at the moment for git-town
	<main>
  \
   1-refactor
     \
      2-bugfix`
	Start_Bugfix_Changes  = "Again, just make a small change just so that we can commit something and create a new PR."
	Start_Bugfix_Finished = "Since git-town is aware that %s is a child of %s, it'll create the PR against the parent, %s\n"

)
