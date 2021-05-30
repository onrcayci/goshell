package goshell

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"text/scanner"
)

func Shell() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf(">> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		argc, argv := parseInput(input)
		interpreter(argc, argv)
	}
}

func parseInput(input string) (int, []string) {
	var args []string
	var scan scanner.Scanner
	scan.Init(strings.NewReader(input))
	for token := scan.Scan(); token != scanner.EOF; token = scan.Scan() {
		arg := scan.TokenText()
		args = append(args, arg)
	}
	return len(args), args
}

func interpreter(argc int, argv []string) {
	if argc == 0 {
		return
	} else {
		switch argv[0] {
		case "help":
			help()
		case "exit":
			exit()
		}
	}
}

func help() {
	helpText := `Go Shell v0.0.1
Available Commands:
help		Prints out help text
exit		Exits the shell
`
	fmt.Println(helpText)
}

func exit() {
	os.Exit(0)
}
