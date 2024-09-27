# Shello

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

###  Executing Commands

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

### Specifying Execution Directory

These methods can be used to specify the directory when running a comamnd:

- `shello.RunWithDir("COMMAND", "DIR")`
- `shello.RunOutputWithDir("COMMAND", "DIR")`
- `shello.PowershellWithDir("COMMAND", "DIR")`
- `shello.ShellOutputWithDir("COMMAND", "DIR")`

Example:
```go
shello.RunWithDir("ls", "/tmp/")
```

### Running Multiple Commands

```go
out, _, _ := shello.RunOutput(`
  echo "Hello Line 1"
  echo "Hello Line 2"
  echo "Hello Line 3"
`)
```

### Executing PowerShell Commands

```go
shello.PowershellCommand(`Write-Host "hello from powershell"`)
```

### Executing Bash/Shell Commands

```go
shello.ShellCommand(`echo "shell or bash?"`)
```

## License

[Apache License 2.0](LICENSE) 