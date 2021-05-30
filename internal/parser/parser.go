package parser

import (
	"strings"
	"text/scanner"
)

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
