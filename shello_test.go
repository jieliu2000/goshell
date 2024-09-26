package shello_test

import (
	"log"
	"os"
	"runtime"
	"strings"
	"testing"

	"github.com/jieliu2000/shello"
)

func TestRunOutputWithDir(t *testing.T) {
	dir := "" // Use the current directory
	expectedOutput := "Hello, World!"
	command := "echo 'Hello, World!'"
	stdout, stderr, err := shello.RunOutputWithDir(command, dir)
	if err != nil {
		t.Errorf("Error executing command: %v", err)
	}
	if stdout != expectedOutput {
		t.Errorf("Unexpected standard output. Expected: %s, Got: %s", expectedOutput, stdout)
	}
	if stderr != "" {
		t.Errorf("Expected no standard error, but got: %s", stderr)
	}
}

func TestRunOutputWithDir_InvalidCommand(t *testing.T) {
	dir := "" // Use the current directory
	command := "invalid_command"
	stdout, stderr, err := shello.RunOutputWithDir(command, dir)
	if err == nil {
		t.Errorf("Expected an error for an invalid command, but got nil")
	}
	if stdout != "" {
		t.Errorf("Expected no standard output for an error, but got: %s", stdout)
	}
	if stderr == "" {
		t.Errorf("Expected standard error for an invalid command, but got empty")
	}
}

func TestRunOutputWithDir_EmptyDirectory(t *testing.T) {
	dir := ""
	expectedOutput := "Hello, World!"
	command := "echo 'Hello, World!'"
	stdout, stderr, err := shello.RunOutputWithDir(command, dir)
	if err != nil {
		t.Errorf("Error executing command: %v", err)
	}
	if stdout != expectedOutput {
		t.Errorf("Unexpected standard output. Expected: %s, Got: %s", expectedOutput, stdout)
	}
	if stderr != "" {
		t.Errorf("Expected no standard error, but got: %s", stderr)
	}
}

func TestRunOutputWithDir_DirectoryNotExists(t *testing.T) {
	dir := "/non-existent-directory"
	expectedOutput := ""
	command := "echo 'Hello, World!'"
	stdout, stderr, err := shello.RunOutputWithDir(command, dir)
	if err == nil {
		t.Errorf("Expected an error for a non-existent directory, but got nil")
	}
	if stdout != expectedOutput {
		t.Errorf("Unexpected standard output. Expected: %s, Got: %s", expectedOutput, stdout)
	}
	if stderr != "" {
		t.Errorf("Expected no standard error for an error, but got: %s", stderr)
	}
}

func TestRunOutputWithDir_1(t *testing.T) {
	dir, err := os.MkdirTemp("", "shello_test")
	if err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll(dir) // clean up

	out, _, err := shello.RunOutputWithDir("pwd", dir+"/")

	if out == "" || !strings.ContainsAny(out, dir) {
		t.Error("output error. didn't find any target directory")
	}
	if err != nil {
		t.Error("error", err)
	}
}

func TestRunOutputOnWindows(t *testing.T) {
	if runtime.GOOS != "windows" {
		t.Skip("These tests are for Windows only")
	}

	for _, tc := range []struct {
		command       string
		expected      string
		expectedError bool
	}{
		{"echo Hello", "Hello", false},
		{"dir non_existent_file.txt", "", true},
	} {
		out, _, err := shello.RunOutput(tc.command)
		if err != nil && !tc.expectedError {
			t.Errorf("RunOutput(%q) returned error: %v", tc.command, err)
		}
		if out != tc.expected {
			t.Errorf("RunOutput(%q) = %q, want %q", tc.command, out, tc.expected)
		}
	}
}

func TestRunOutputOnNonWindows(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("These tests are for non-Windows only")
	}

	for _, tc := range []struct {
		command       string
		expected      string
		expectedError bool
	}{
		{"echo Hello", "Hello", false},
		{"ls non_existent_file.txt", "", true},
	} {
		out, _, err := shello.RunOutput(tc.command)
		if err != nil && !tc.expectedError {
			t.Errorf("RunOutput(%q) returned error: %v", tc.command, err)
		}
		if out != tc.expected {
			t.Errorf("RunOutput(%q) = %q, want %q", tc.command, out, tc.expected)
		}
	}
}

func TestRunOutput(t *testing.T) {
	for _, tc := range []struct {
		command       string
		expected      string
		expectedError bool
	}{
		{"echo Hello", "Hello", false},
		{"dir non_existent_file.txt", "", true},
	} {
		out, _, err := shello.RunOutput(tc.command)
		if err != nil && !tc.expectedError {
			t.Errorf("RunOutput(%q) returned error: %v", tc.command, err)
		}
		if out != tc.expected {
			t.Errorf("RunOutput(%q) = %q, want %q", tc.command, out, tc.expected)
		}
	}
}

func TestPowershellOutputWithDir(t *testing.T) {
	if runtime.GOOS != "windows" {
		t.Skip("These tests are for Windows only")
	}
	// Test PowershellOutputWithDir method with valid directory
	if _, _, err := shello.PowershellOutputWithDir("Write-Host Hello from PowershellOutputWithDir", "."); err != nil {
		t.Errorf("expected no error, got error: %v", err)
	}

	// Test PowershellOutputWithDir method with invalid directory
	if _, _, err := shello.PowershellOutputWithDir("Write-Host Hello from PowershellOutputWithDir", "/invalid-directory"); err == nil {
		t.Error("expected error because dir is invalid, got no error")
	}
}

func TestRun(t *testing.T) {
	shello.Run("echo Hello from Run")
}

func TestRunMultipleLine(t *testing.T) {
	out, _, _ := shello.RunOutput("echo Hello Line 1\necho Hello Line 2\necho Hello Line 3 ")
	if !strings.Contains(out, "Hello Line 1") {
		t.Error("output error")
	}
}

func TestRunWithDir(t *testing.T) {
	shello.RunWithDir("echo Hello from RunWithDir", ".")
}

func TestShellOutputWithDir(t *testing.T) {
	// Test ShellOutputWithDir method with valid directory
	if _, _, err := shello.ShellOutputWithDir("echo Hello from ShellOutputWithDir", "."); err != nil {
		t.Errorf("expected no error, got error: %v", err)
	}

	// Test ShellOutputWithDir method with invalid directory
	if _, _, err := shello.ShellOutputWithDir("echo Hello from ShellOutputWithDir", "/invalid-directory"); err == nil {
		t.Errorf("expected error because dir is invalid, got no error")
	}
}

func TestExecWithInvalidCommand(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("These tests are for non-Windows only")
	}
	// Test 1: Verify execution error with invalid command
	out, errOut, err := shello.Exec("bash", "invalid_command", "")
	if err == nil {
		t.Error("Expected error for invalid command")
	}
	if out != "" {
		t.Errorf("Unexpected output for invalid command: %s", out)
	}
	if errOut == "" {
		t.Errorf("Unexpected error output for invalid command: %s", errOut)
	}
}

func TestExecWithCommandAndOutput(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("These tests are for non-Windows only")
	}
	shello.TrimOutput = false
	// Test 1: Verify successful execution with command and expected output
	out, errOut, err := shello.Exec("bash", `echo 'This is the expected output'`, "")
	if err != nil {
		t.Errorf("Error executing command: %v", err)
	}
	expectedOut := "This is the expected output\n"
	if out != expectedOut {
		t.Errorf("Unexpected output. Expected: %s, Actual: %s", expectedOut, out)
	}
	if errOut != "" {
		t.Errorf("Unexpected error output: %s", errOut)
	}
	shello.TrimOutput = true
}

func TestExec(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("These tests are for non-Windows only")
	}
	shello.TrimOutput = false
	// Test 1: Verify successful execution with correct command
	out, errOut, err := shello.Exec("bash", "echo 'Hello, World!'", "")
	if err != nil {
		t.Errorf("Error executing command: %v", err)
	}
	if out != "Hello, World!\n" {
		t.Errorf("Unexpected output. Expected: 'Hello, World!\n', Actual: %s", out)
	}
	if errOut != "" {
		t.Errorf("Unexpected error output: %s", errOut)
	}

	// Test 2: Verify execution error with incorrect command
	out, errOut, err = shello.Exec("bash", "invalid_command", "")
	if err == nil {
		t.Error("Expected error for invalid command")
	}
	if out != "" {
		t.Errorf("Unexpected output for invalid command: %s", out)
	}
	if errOut == "" {
		t.Error("Error output should not be empty since this is an invalid command")
	}

	// Test 3: Verify execution with multiple commands
	out, errOut, err = shello.Exec("bash", "echo 'First Command'; echo 'Second Command';", "")
	if err != nil {
		t.Errorf("Error executing commands: %v", err)
	}
	expectedOut := "First Command\nSecond Command\n"
	if out != expectedOut {
		t.Errorf("Unexpected output. Expected: %s, Actual: %s", expectedOut, out)
	}
	if errOut != "" {
		t.Errorf("Unexpected error output: %s", errOut)
	}

	// Test 4: Verify execution with special characters in command
	out, errOut, err = shello.Exec("bash", `echo '"Hello, World!"'`, "")
	if err != nil {
		t.Errorf("Error executing command: %v", err)
	}
	expectedOut = "\"Hello, World!\"\n"
	if out != expectedOut {
		t.Errorf("Unexpected output. Expected: %s, Actual: %s", expectedOut, out)
	}
	if errOut != "" {
		t.Errorf("Unexpected error output: %s", errOut)
	}

	shello.TrimOutput = true

}

func TestExecWithInvalidShell(t *testing.T) {
	// Test 1: Verify execution error with invalid shell
	out, errOut, err := shello.Exec("invalid_shell", "echo 'Hello, World!'", "")
	if err == nil {
		t.Error("Expected error for invalid shell")
	}
	if out != "" {
		t.Errorf("Unexpected output for invalid shell: %s", out)
	}
	if errOut != "" {
		t.Errorf("Unexpected error output for invalid shell: %s", errOut)
	}
}
