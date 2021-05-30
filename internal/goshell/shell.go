package goshell

import (
	"bufio"
	"fmt"
	"os"

	"github.com/onrcayci/goshell/internal/interpreter"
	"github.com/onrcayci/goshell/internal/parser"
)

func Shell() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf(">> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		argc, argv := parser.ParseInput(input)
		interpreter.Interpreter(argc, argv)
	}
}
