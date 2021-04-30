package main

import (
	"bufio"
	"os"

	"github.com/onrcayci/goshell/internal/goshell"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		scanner.Scan()
		input := scanner.Text()
		argc, argv := goshell.ParseInput(input)
		goshell.Interpreter(argc, argv)
	}
}
