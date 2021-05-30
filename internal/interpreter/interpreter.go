package interpreter

import "github.com/onrcayci/goshell/internal/command"

func Interpreter(argc int, argv []string) {
	if argc == 0 {
		return
	} else {
		switch argv[0] {
		case "help":
			command.Help()
		case "exit":
			command.Exit()
		}
	}
}
