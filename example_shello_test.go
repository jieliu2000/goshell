package shello_test

import (
	"fmt"
	"log"

	"github.com/jieliu2000/shello"
)

func ExampleRunOutput() {
	fmt.Println("Example of echo 𝜋:")

	// run a command with output
	out, errout, err := shello.RunOutput("echo 𝜋")
	if err != nil {
		log.Printf("error: %v\n", err)
		fmt.Print(errout)
		return
	}
	// The default value of shello.TrimOutput is true, so the output will be trimmed and the trailing newline will be removed.
	fmt.Print(string(out))

	// Because the last output is not a newline, we need to add it here before printing.
	fmt.Println("\nExample of run without TrimOutput:")
	// run a command without trim output
	shello.TrimOutput = false
	out, errout, err = shello.RunOutput("echo 𝜋")
	fmt.Print(string(out))
	// Output:
	// Example of echo 𝜋:
	// 𝜋
	// Example of run without TrimOutput:
	// 𝜋

}
