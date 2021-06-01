package goshell

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/onrcayci/goshell/internal/interpreter"
	"github.com/onrcayci/goshell/internal/parser"
)

// Shell implements the main shell infinite loop. It accepts input from the user,
// parses the input into tokens and then uses the interpreter function to execute
// the desired command.
func Shell() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf(">> ")
		input, err := reader.ReadString('\n')
		if err == io.EOF {
			os.Exit(1)
		} else if err != nil {
			panic(err)
		}
		argc, argv := parser.ParseInput(input)
		interpreter.Interpreter(argc, argv)
	}
}
