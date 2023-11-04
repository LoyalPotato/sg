# TODO

## v1

For the first release, what are the tasks that need to be done.

- [x] Make sure git-town is working
I can't work with the go package. It doesn't seem like it was made for that :(

- [x] Test how to have running interactivity
- [x] Test if it's running in a git repository
  - [x] Suggest `git init` or wtv the command is
  - [x] This should be done for all commands
- [x] Get the setup command working
- [x] Setup Help command
- [x] In setup, ask for what the main branch is. Gittown needs this configured else it'll fail when running the cmd in the cli?
- [ ] Add goreleaser

### Later

- [ ] Man pages, markdown docs [cobra](https://github.com/spf13/cobra/blob/main/site/content/docgen/_index.md)

## v2

- [ ] [Shell completions](https://github.com/spf13/cobra/blob/main/site/content/completions/_index.md)
- [ ] State persistance for continue. Git town keeps its own state and has an interpreter for that o.o Sounds complicated, but it's all just a matter of breaking down and time :)
