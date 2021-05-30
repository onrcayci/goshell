package interpreter

import (
	"fmt"

	"github.com/onrcayci/goshell/internal/command"
)

func Interpreter(argc int, argv []string) {
	if argc == 0 {
		return
	} else {
		switch argv[0] {
		case "help":
			command.Help()
		case "quit":
			command.Quit()
		case "set":
			err := command.Set(argc, argv)
			if err != nil {
				fmt.Println(err.Error())
			}
		case "print":
			err := command.Print(argc, argv)
			if err != nil {
				fmt.Println(err.Error())
			}
		default:
			fmt.Printf("%s: command not found\n", argv[0])
		}
	}
}
