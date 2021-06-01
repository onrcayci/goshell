package interpreter

import (
	"github.com/onrcayci/goshell/internal/command"
)

// Interpreter takes in the tokenized input slice (argv) and the length of the slice(arc)
// and uses a switch statement to determine which shell command to execute.
// Supported commands: "help", "quit", "set", "print" and "run".
func Interpreter(argc int, argv []string) {
	if argc == 0 {
		return
	} else {
		command.Execute(argc, argv)
	}
}
