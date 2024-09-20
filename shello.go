package shello

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"runtime"
	"strings"
)

// `ShellOutput` returns the output of shell command, and any errors.
func ShellOutput(command string) (string, string, error) {
	return ShellOutputWithDir(command, "")
}

func ShellOutputWithDir(command, dir string) (string, string, error) {
	return Exec("bash", command, dir)
}

// `ShellCommand` executes the shell command.
func ShellCommand(command string) {
	out, errout, err := ShellOutput(command)

	if err != nil {
		log.Printf("error: %v\n", err)
		fmt.Print(errout)
	}
	fmt.Print(out)
}

// `PowershellOutput` returns the output of powershell command, and any errors.
func PowershellOutput(command string) (string, string, error) {
	return PowershellOutputWithDir(command, "")
}

func PowershellOutputWithDir(command, dir string) (string, string, error) {
	return Exec("powershell.exe", command, dir)
}

// `PowershellCommand` executes the powershell command.
func PowershellCommand(command string) {
	out, errout, err := PowershellOutput(command)

	if err != nil {
		log.Printf("error: %v\n", err)
		fmt.Print(errout)
	}

	fmt.Print(out)
}

// `Exec` just exectes the command
func Exec(shell, command, dir string) (string, string, error) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	cmd := exec.Command(shell, "-c", command)

	if dir != "" {
		cmd.Dir = dir
	}
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()

	return outputString(&stdout), outputString(&stderr), err
}

// `Run` executes the same command for shell and powershell
func Run(cmd string) {
	RunWithDir(cmd, "")
}

// RunWithDir runs a command in a specified directory. This function determines the operating system and calls the appropriate function to execute the command.
// If an error occurs during execution, the error message is logged and the error output is printed.
// Parameters:
// - cmd string: The command to be executed.
// - dir string: The directory in which the command will be executed.
// This function determines the operating system and calls the appropriate function to execute the command.
// If an error occurs during execution, the error message is logged and the error output is printed.
// Finally, the command output is printed.
func RunWithDir(cmd, dir string) {
	var (
		out    string
		errout string
		err    error
	)

	if runtime.GOOS == "windows" {
		out, errout, err = PowershellOutputWithDir(cmd, dir)
	} else {
		out, errout, err = ShellOutputWithDir(cmd, dir)
	}

	if err != nil {
		log.Printf("error: %v\n", err)
		fmt.Print(errout)
	}

	fmt.Print(out)
}

// `RunOutput` returns the output of the shared command for shell and powershell
func RunOutput(command string) (string, string, error) {
	return RunOutputWithDir(command, "")
}

// RunOutputWithDir runs a command in a specified directory and returns its standard output, standard error, and any error encountered.
// Parameters:
// - command string: The command to be executed.
// - dir string: The directory in which the command should be executed. If empty, the current directory is used.
// Return values:
// - string: The standard output of the command.
// - string: The standard error of the command.
// - error: If an error occurs, the corresponding error message is returned.
func RunOutputWithDir(command, dir string) (string, string, error) {

	if runtime.GOOS == "windows" {
		return PowershellOutputWithDir(command, dir)
	} else {
		return ShellOutputWithDir(command, dir)
	}
}

// outputString returns the string representation of a bytes.Buffer.
// If the buffer is nil, an empty string is returned.
// The output is trimmed if `TrimOutput` is set to true.
func outputString(buf *bytes.Buffer) string {
	if buf == nil {
		return ""
	}
	if TrimOutput {
		return strings.TrimSpace(buf.String())
	}
	return buf.String()
}

// `TrimOutput` is a flag to trim the output of commands
var TrimOutput = true
