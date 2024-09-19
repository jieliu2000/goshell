package gosh

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"os/exec"
	"runtime"
)

// `ShellOutput` returns the output of shell command, and any errors.
func ShellOutput(command string) (string, string, error) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	cmd := exec.Command("bash", "-c", command)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()

	return stdout.String(), stderr.String(), err
}

func ShellOutputWithDir(command, dir string) (string, string, error) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	cmd := exec.Command("bash", "-c", command)
	cmd.Dir = dir
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()

	return stdout.String(), stderr.String(), err
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
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	cmd := exec.Command("powershell.exe", command)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()

	return stdout.String(), stderr.String(), err
}

func PowershellOutputWithDir(command, dir string) (error, string, string) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	cmd := exec.Command("powershell.exe", command)
	cmd.Dir = dir
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()

	return err, stdout.String(), stderr.String()
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
func Exec(shell, command string) (string, string, error) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	cmd := exec.Command(shell, "-c", command)

	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()

	return stdout.String(), stderr.String(), err
}

// `Run` executes the same command for shell and powershell
func Run(cmd string) {
	var (
		out    string
		errout string
		err    error
	)

	if runtime.GOOS == "windows" {
		out, errout, err = PowershellOutput(cmd)
	} else {
		out, errout, err = ShellOutput(cmd)
	}

	if err != nil {
		log.Printf("error: %v\n", err)
		fmt.Print(errout)
	}

	fmt.Print(out)
}

// `RunOutput` returns the output of the shared command for shell and powershell
func RunOutput(command string) (string, string, error) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	var cmd *exec.Cmd

	if runtime.GOOS == "windows" {
		cmd = exec.Command("powershell.exe", command)
	} else {
		cmd = exec.Command("bash", "-c", command)
	}
	if cmd == nil {
		return "", "", errors.New("command is nil")
	}
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()

	return stdout.String(), stderr.String(), err
}
