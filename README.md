# go-setup

CLI tool to setup Golang project.

## Installation

Execute `go install github.com/dark-shade/go-setup`.

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

Global Flags:
      --config string   config file (default is $HOME/.go-setup.yaml)
```


