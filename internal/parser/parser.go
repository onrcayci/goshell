package parser

import (
	"strings"
	"text/scanner"
)

// Function to parse the input into string tokens.
// Retunrs the number of tokens (argc) and the tokens (args).
func ParseInput(input string) (int, []string) {
	var args []string
	var scan scanner.Scanner
	scan.Init(strings.NewReader(input))
	for token := scan.Scan(); token != scanner.EOF; token = scan.Scan() {
		arg := scan.TokenText()
		args = append(args, arg)
	}
	return len(args), args
}
