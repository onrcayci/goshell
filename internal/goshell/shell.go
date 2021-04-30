package goshell

import (
	"fmt"
	"os"
	"strings"
)

func ParseInput(input string) (int, []string) {
	input = strings.TrimSuffix(input, "\r\n")
	args := strings.Split(input, " ")
	return len(args), args
}

func Interpreter(argc int, argv []string) {
	if argc == 0 {
		return
	} else {
		switch argv[0] {
		case "help":
			Help()
		case "exit":
			Exit()
		}
	}
}

func Help() {
	helpText := `Go Shell v0.0.1
Available Commands:
help		Prints out help text
exit		Exits the shell
`
	fmt.Println(helpText)
}

func Exit() {
	os.Exit(0)
}
