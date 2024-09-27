# shello

[![GoDoc](https://godoc.org/github.com/jieliu2000/shello?status.svg)](https://pkg.go.dev/github.com/jieliu2000/shello)
[![Go Report Card](https://goreportcard.com/badge/github.com/jieliu2000/shello)](https://goreportcard.com/jieliu2000/shello)



> A golang library for executing bash & powershell commands easily.

shello is a golang library for executing bash & powershell commands easily. It is a folked version of [gosh](https://github.com/abdfnx/gosh) by [abdfnx](https://github.com/abdfnx). I folked gosh with a new name because it seems that abdfnx is not maintaining gosh anymore.

**Please note that shello's API is NOT compatible with the original gosh API**.

## Install

```bash
go get -u github.com/jieliu2000/shello
```

## Global Variables

Shello has a global variable `shello.TrimOutput` which is used to trim the output of a command. By default, it is set to `true`, which means the output will be trimmed. You can set it to `false` if you want to keep the original output of a command.

```go
shello.TrimOutput = false
```

## Examples

### Run command

These methods can be used to run a command without any returned output or error:
- `shello.Run("COMMAND")`
- `shello.ShellCommand("COMMAND")`
- `shello.PowershellCommand("COMMAND")`

These methods can be used to run a command with output:
- `shello.RunOutput("COMMAND")`
- `shello.ShellOutput("COMMAND")`
- `shello.PowershellOutput("COMMAND")`

Shello will automatically detect the OS and run the command for either bash or powershell if you use `shello.Run("COMMAND")` or `shello.RunOutput("COMMAND")`.

These method will execute the command in the current directory by default. If you want to specify the directory for a command, see [Specify the directory for a command](#speicify-the-directory-for-a-command) below.

```go
package main

import (
  "fmt"
  "log"

  "github.com/jieliu2000/shello"
)

// run a command
shello.Run("git status")

// run a command with output
out, errout, err := shello.RunOutput("echo 𝜋")

if err != nil {
  log.Printf("error: %v\n", err)
  fmt.Print(errout)
}

fmt.Print(out)
```

The xxxOutput methods will return three values: the output of the command, the error output and an error.

### Speicify the directory for a command

These methods can be used to specify the directory when running a comamnd:

- `shello.RunWithDir("COMMAND", "DIR")`
- `shello.RunOutputWithDir("COMMAND", "DIR")`
- `shello.PowershellWithDir("COMMAND", "DIR")`
- `shello.ShellOutputWithDir("COMMAND", "DIR")`

Example:
```go
shello.RunWithDir("ls", "/tmp/")
```

### Run mulitple commands together

You can simple pass a mulitple-line string to the shello methods to run mulitple commands together.
```go
	out, _, _ := shello.RunOutput(`echo "Hello Line 1"
  echo "Hello Line 2"
  echo "Hello Line 3"`)

```

### Run Powershell Command(s)

```go
package main

import (
  "fmt"
  "log"

  "github.com/jieliu2000/shello"
)

// run a command
shello.PowershellCommand(`Write-Host "hello from powershell"`)

// run a script
shello.PowershellCommand(`
  $git_username = git config user.name

  Write-Host $git_username
`)

// run a command with output
out, errout, err := shello.PowershellOutput(`[System.Environment]::SetEnvironmentVariable("Path", $Env:Path + ";$APP_PATH\bin", [System.EnvironmentVariableTarget]::User)`)

if err != nil {
  log.Printf("error: %v\n", err)
  fmt.Print(errout)
}

fmt.Print(out)
```

### Run Bash/Shell Command(s)

```go
package main

import (
  "fmt"
  "log"

  "github.com/jieliu2000/shello"
)

// run a command
shello.ShellCommand(`echo "shell or bash?"`)

// run a script
shello.ShellCommand(`
  mood="👨‍💻"

  if [ $mood != "😪" ]; then
    echo "still coding"
  fi
`)

// run a command with output
out, errout, err := shello.ShellOutput(`curl --silent "https://get-latest.onrender.com/docker/compose"`)

if err != nil {
  log.Printf("error: %v\n", err)
  fmt.Print(errout)
}

fmt.Print(out)
```
