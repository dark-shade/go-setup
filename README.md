![codeql workflow](https://github.com/dark-shade/go-setup/actions/workflows/codeql-analysis.yml/badge.svg)
![build workflow](https://github.com/dark-shade/go-setup/actions/workflows/build-across-matrix.yml/badge.svg)

# go-setup

CLI tool to setup Golang project. It loosely follows [golang-standards/project-layout](https://github.com/golang-standards/project-layout).

## Installation

Execute `go install github.com/dark-shade/go-setup@latest`.

## Usage

```bash
$ go-setup --help
A CLI app that provides the ability to setup and modify structure of multiple types of golang projects.
It loosely follows https://github.com/golang-standards/project-layout

Usage:
  go-setup [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  init        Initializes a project

Flags:
      --config string   config file (default is $HOME/.go-setup.yaml)
  -h, --help            help for go-setup
  -t, --toggle          Help message for toggle

Use "go-setup [command] --help" for more information about a command.
```

```bash
$ go-setup init --help
Initializes a project by adding recommended directory structure and files.

Usage:
  go-setup init [flags]

Flags:
  -a, --author string         author name and email, e.g. Jane Doe jane.doe@gmail.com
  -f, --full                  initializes all files and directories in the recommend layout
  -h, --help                  help for init
  -i, --license string        initializes the license (default "mit")
  -l, --location string       location for project structure setup (default ".")
  -m, --moduleP-path string   module path for go mod init (default ".")
  -o, --ops                   initializes all the operations related files
  -p, --profile strings       profile to use for project setup (default [default])

Global Flags:
      --config string   config file (default is $HOME/.go-setup.yaml)
```

### Usage of Profiles

Profiles are special files and directories that a user wants to add during project setup which are not covered by [golang-standards/project-layout](https://github.com/golang-standards/project-layout). User has the ability to add custom profiles which when specified using the `go-setup init -p <profile-names>` will add the files present in the profiles to the target location.

In order to use profiles with go-setup user needs to first execute `go-setup init -c` which will create a directory on path `$HOME/.go-setup/profiles` (if not already present, NOTE: any `go init` execution will create this directory if it is not already present ). Then to add a profile create a directory with the profile name as the directory name under `$HOME/.go-setup/profiles` (for example, say you want a profile named `templates` then create `$HOME/.go-setup/profiles/templates`). Add all the files and directories in the profile directory. Then to setup your projects using profile execute `go-setup init -p <profile-names>`, NOTE: this will also create other bare files and directories if not present in the profile.

#### Example of profile usage

In this example we will use two profiles `templates` and `js`. The `template` profile will contain `tmpl1.html` and `tmpl2.html` files and `static` directory. The `js` profile will contain `main.js`. Steps to do this:
1. Check if `$HOME/.go-setup/profiles` exists, if not execute `go-setup init -c`.
2. Execute `mkdir $HOME/.go-setup/profiles/templates` and `mkdir $HOME/.go-setup/profiles/js`.
3. Put `tmpl1.html`, `tmpl2.html` and `static` directory in `$HOME/.go-setup/profiles/templates`.
4. Put `main.js` in `$HOME/.go-setup/profiles/js`.
5. Execute `go-setup init -p templates,js -l repos/project-repo`, in this the `-l` flag specifies where to setup project (in this case `repos/project-repo`). **NOTE:** profiles are processed in sequence and the files and directories are not overwritten so if `template` and `js` profiles contain the same file then the file present in `template` which was created first will be final.
