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
	GitTown_Sync   = "git-town sync"

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
	Start_Hack = `Let's create a new branch from %s to start working on the cleanup.
This can be done with the following command from git-town:
%s
Type it out below to run it :)
`
	Start_Hack_Finished = `Well done!
This command just created your new feature branch from the %s branch you setup and checked it out.
`

	Start_Refactor_Changes = `Now make some changes to any file.
Just something simple like creating a new file and writing "Refactoring", for example.
Once you're finished, proceed with the guide.`
	Start_Refactor_Finished = `Now all that's left for this step is to finish creating the PR, through the link opened.
If the link didn't open in your default browser, manually open it to conclude the creation of the PR.

And we've successfully put our first changes for review!`

	Start_Bugfix = `We can now start working on the bugfix on top of these changes while they are still in review, one of the advantages of the stacking workflow!
To build on top of the previous changes, in the refactor branch still, we use the following git-town command:
%s
Give it a go:`
	Start_Bugfix_Tree = `This is what the branch hierarchy is at the moment for git-town
%s
  \
   1-refactor
     \
      2-bugfix

`
	Start_Bugfix_Changes  = "Again, just make a small change just so that we can commit something and create a new PR."
	Start_Bugfix_Finished = "Since git-town is aware that %s is a child of %s, it'll create the PR against the parent, %s.\n"

	Start_Feature = `
Finally, we can begin working on our feature.
Like the previous branch we want to create it on top of all of the changes (refactor and bugfix).

You know the drill ;)
%s`
	Start_Feature_Tree = `The hierarchy now looks like this
%s
  \
   1-refactor
     \
      2-bugfix
        \
         3-feature

`
	Start_Feature_Refactor         = "The feature branch was created successfully but we've got some comments on %s's PR. Let's apply them before working on the feature.\n"
	Start_Feature_Refactor_Changes = "We've checked %s out. Create a new file and write \"Reviewer note\" inside.\n"
	Start_Feature_Sync             = `With the changes made we can now go back to %s.
But now we don't have the fixes that we just made.
Luckily git-town facilitates synchronizing all our child branches, allowing us to avoid the tedious manual synchronization.
This is achieved by running this simple command:
%s

What it'll do is go on each child branch (%s and %s) and merge the changes.
`
	Start_Feature_Changes  = "We're almost finished with making changes. Do it one more time for in this feature branch."
	Start_Feature_Finished = "We've finished working on the feature so, like the previous two, we'll create a new PR."

	Start_Ship = `
We got the final approve for our refactor!

Let's merge it into %s. We can do it with the %s command in git-town:
%s
`
	Start_Finished = `And that's it!

Stacking workflow allows you to separate PRs that adhere to the single responsability principle, leading to smaller PRs.
It also allows you to work faster since you can continue chaining features while waiting for review/approval.

This workflow can be achieved with just the commands of git, but git-town is a nice wrapper around git with some extra functionalities that make it easier to follow this workflow, like the %s command.

You can find resources in the repository of this guide: %s

Thank you for using this guide and I hope it helped you in understanding the stacking workflow and how to use git-town to that effect.
At the moment the guide doesn't allow reruns :( since it's using hard coded values. If you want to run it again, just delete all the branches created and run again from start.

Best of luck on your stacking :)
`
)
