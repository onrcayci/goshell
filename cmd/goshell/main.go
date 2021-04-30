package main

import (
	"bufio"
	"os"

	"github.com/onrcayci/goshell/internal/goshell"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		input, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		argc, argv := goshell.ParseInput(input)
		goshell.Interpreter(argc, argv)
	}
}
