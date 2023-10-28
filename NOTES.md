# Planning

Like I said in the README, I want this to be an interactive experience.

I'm not sure, but I was thinking of using **git-town** directly, to avoid any extra work in showing this workflow.
I believe it has a go package, so could I make calls to the api through that?

- Wrap around git-town commands and guide the user in using them.
- User should run this guide in an example repository, where it'll create the branches, mr and whatnot
- Handle misinputs?
- Have an easter egg command :)
- Autocomplete when user is typing? (doubt?)
- After installation, show what command to start (use "start"?)
- Display help text when no options are passed or the regular -h --help

## Commands

- help
- setup
- start
- wherewasi (shows last executed command)
- continue
- uninstall
  - Show how? Or if possible remove itself? I've seen people run a script with curl and such

## Setup

What do we need to have everything working before starting?

- be in a git environment
- git-town setup

## Start

Don't start unless setup as been completed.
Check if the user has already started and warn if they want to start over. [y/n]

## In Prog

- Should the setup be allowed to be run again?
- Handle errors, signals and whatnot
- Should it be like a state machine? So that the user can continue if they stop? How could we store this information? Git-town seems to be able to store it somehow, because of the `git-town continue`

## Reset/restart

> Standard in/out/err, signals, exit codes and other mechanisms ensure that different programs click together nicely.

## Notes

> Discoverable CLIs have comprehensive help texts, provide lots of examples, suggest what command to run next, suggest what to do when there is an error. There are lots of ideas that can be stolen from GUIs to make CLIs easier to learn and use, even for power users.

This is what I want mine to feel like!
> Ironically, though, the CLI has embodied an accidental metaphor all along: it’s a conversation.
> Running one command to set up a tool and then learning what commands to run to actually start using it.
> It means giving the user the feeling that you are on their side, that you want them to succeed, that you have thought carefully about their problems and how to solve them.
> redirection of stderr to /dev/null

### Help

>The concise help text should only include:
> A description of what your program does.
> One or two example invocations.
> Descriptions of flags, unless there are lots of them.
> An instruction to pass the --help flag for more information.
> Use color with intention.

- Provide a support path for feedback and issues. A website or GitHub link in the top-level help text is common.

> However, not everyone knows about man, and it doesn’t run on all platforms, so you should also make sure your terminal docs are accessible via your tool itself. For example, git and npm make their man pages accessible via the help subcommand, so npm help ls is equivalent to man npm-ls.

- Check for `NO_COLOR` in env [no color](https://no-color.org/)
The NO_COLOR environment variable is set.

[Active Help - Cobra](https://github.com/spf13/cobra/blob/main/site/content/active_help.md) Seems like an interesting thing.
