# shello

> A golang library for executing bash & powershell commands easily.

shello is a golang library for executing bash & powershell commands easily. It is a folked version of [gosh](https://github.com/abdfnx/gosh) by [abdfnx](https://github.com/abdfnx). I folked gosh with a new name because it seems that abdfnx is not maintaining gosh anymore.

**Please note that shello's API is NOT compatible with the original gosh API**.

## Install

```bash
go get -u github.com/jieliu2000/shello
```

## Examples

### Run one command on both shell and powershell

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
err, out, errout := shello.RunOutput("echo 𝜋")

if err != nil {
  log.Printf("error: %v\n", err)
  fmt.Print(errout)
}

fmt.Print(out)
```

### How `shello.Run("COMMAND")` works ?

```go
// `Run` executes the same command for shell and powershell
func Run(cmd string) {
	err, out, errout := ShellOutput("")

	if runtime.GOOS == "windows" {
		err, out, errout = PowershellOutput(cmd)
	} else {
		err, out, errout = ShellOutput(cmd)
	}

	if err != nil {
		log.Printf("error: %v\n", err)
		fmt.Print(errout)
	}

	fmt.Print(out)
}
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
