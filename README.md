# README

This project is an example go cli.

## Commands

It implements the following commands:

- `echo`: takes a single parameter, and echos back the parameter in the output
- `version`: echos back the version of the cli executable

## Github actions

- Build and Release: compile and build the executable, then create release executables for the below mentioned platforms using [GoReleaser](https://goreleaser.com), as well as the homebrew release using the [GoReleaser Homebrew Taps feature](https://goreleaser.com/customization/homebrew/?h=homebrew).  The homebrew release repo is at https://github.com/spinguard/homebrew-example-cli.git.

- Test: test the executable by running the `version` command and asserting the correct release version

## Platforms supported

- Mac Silicon
- Mac Intel
- Linux ARM
- Linux Intel
