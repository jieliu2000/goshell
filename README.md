# shello

[![GoDoc](https://godoc.org/github.com/jieliu2000/shello?status.svg)](https://pkg.go.dev/github.com/jieliu2000/shello)
[![Go Report Card](https://goreportcard.com/badge/github.com/jieliu2000/shello)](https://goreportcard.com/report/github.com/jieliu2000/shello)

> A golang library for executing bash & powershell commands easily.


Shello is a Go library designed to make executing shell (bash) and PowerShell commands straightforward. It is a fork of [gosh](https://github.com/abdfnx/gosh) by [abdfnx](https://github.com/abdfnx), tailored with improvements and maintained for ongoing support.

**Note:** Shello's API is not compatible with the original gosh API.


## Install

```bash
go get -u github.com/jieliu2000/shello
```

## Key Features
* Execute shell and PowerShell commands with ease.
* Automatic detection of the operating system to choose between shell or PowerShell.
* Control over trimming of command output.
* Support for running commands in specified directories.
* Execute multiple commands within a single string.

## Configuration

Shello includes a global variable `TrimOutput` to control whether command output should be trimmed. By default, it is set to `true`.

```go
import "github.com/jieliu2000/shello"

shello.TrimOutput = false // Disable trimming.
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

func main() {
  // Execute a command without capturing output.
  shello.Run("git status")

  // Execute a command and capture the output.
  out, errout, err := shello.RunOutput("echo 𝜋")
  if err != nil {
    log.Printf("error: %v\n", err)
    fmt.Print(errout)
  }
  fmt.Print(out)
}
```

### Speicify the directory for command

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
out, _, _ := shello.RunOutput(`
  echo "Hello Line 1"
  echo "Hello Line 2"
  echo "Hello Line 3"
`)
```

### Run Powershell Command(s)

```go
shello.PowershellCommand(`Write-Host "hello from powershell"`)
```

### Run Bash/Shell Command(s)

```go
shello.ShellCommand(`echo "shell or bash?"`)
```
