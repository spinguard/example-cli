# README

This project is an example go cli

## Commands

It implements the following commands:

- `echo`: takes a single parameter, and echos back the parameter in the output
- `version`: echos back the version of the cli executable

## Github actions

- Build and Release: compile and build the executable, then create release executables for the below mentioned platforms
- Test: test the executable by running the `version` command and asserting the correct release version

## Platforms supported

- Mac Silicon
- Mac Intel
- Linux ARM
- Linux Intel

The 
