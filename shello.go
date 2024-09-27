// This file was modified from https://github.com/abdfnx/gosh/blob/trunk/gosh.go. The methods signatures are NOT compatible with the original.
package shello

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"runtime"
	"strings"
)

// ShellOutput function executes the provided shell command and returns the standard output, standard error, and any error encountered.
// Parameters:
// - command string: The shell command to be executed.
// Return values:
// - string: The standard output of the executed command.
// - string: The standard error of the executed command.
// - error: If an error occurs during execution, the corresponding error message is returned.
func ShellOutput(command string) (string, string, error) {
	return ShellOutputWithDir(command, "")
}

// ShellOutputWithDir executes a shell command in a specified directory and returns the standard output, standard error, and any error encountered.
// Parameters:
// - command string: The shell command to be executed.
// - dir string: The directory in which the command should be executed.
// Return values:
// - string: The standard output of the executed command.
// - string: The standard error of the executed command.
// - error: If an error occurs during execution, the corresponding error message is returned.
func ShellOutputWithDir(command, dir string) (string, string, error) {
	return Exec("bash", command, dir)
}

// ShellCommand executes the given shell command and prints the output.
// Parameters:
// - command string: The shell command to be executed.
// return value:
// - None. The function does not return any value.
func ShellCommand(command string) {
	out, errout, err := ShellOutput(command)

	if err != nil {
		log.Printf("error: %v\n", err)
		fmt.Print(errout)
	}
	fmt.Print(out)
}

// PowershellOutput is a function that executes a PowerShell command and returns its output.
// It takes a string 'command' as input which represents the PowerShell command to be executed.
// The function returns two strings: the first string represents the standard output of the command,
// the second string represents the standard error output of the command,
// and an error object which contains any error information that occurred during execution.
func PowershellOutput(command string) (string, string, error) {
	return PowershellOutputWithDir(command, "")
}

// PowershellOutputWithDir function executes a PowerShell command within a specified directory and returns the output.
// Parameters:
// - command string: The PowerShell command to be executed.
// - dir string: The directory in which the command will be executed.
// Return values:
// - string: The standard output of the executed command.
// - string: The standard error output of the executed command.
// - error: If an error occurs during the execution, the corresponding error message is returned.
func PowershellOutputWithDir(command, dir string) (string, string, error) {
	return Exec("powershell.exe", command, dir)
}

// PowershellCommand executes a given Powershell command and prints the output.
// Parameters:
// - command string: The Powershell command to be executed.
// return value:
// - None. The function does not return any value.
func PowershellCommand(command string) {
	out, errout, err := PowershellOutput(command)

	if err != nil {
		log.Printf("error: %v\n", err)
		fmt.Print(errout)
	}

	fmt.Print(out)
}

// Exec function executes a shell command and returns the standard output, standard error, and any error that occurred.
// Parameters:
// - shell string: The shell to be used for executing the command.
// - command string: The command to be executed.
// - dir string: The directory in which the command will be executed. If empty, the current directory will be used.
// Return values:
// - string: The standard output of the executed command.
// - string: The standard error of the executed command.
// - error: If an error occurs during the execution of the command, the corresponding error message is returned.
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

// Run function executes the given command.
// Parameters:
// - cmd string: The command to be executed.
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

// RunOutput runs the given command and returns its standard output, standard error, and any error encountered.
// If the command is successful, the error will be nil.
// Parameters:
// - command string: The command to be executed.
// Return values:
// - string: The standard output of the command.
// - string: The standard error of the command.
// - error: If an error occurs, the corresponding error message is returned.
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
