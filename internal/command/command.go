package command

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/onrcayci/goshell/internal/memory"
	"github.com/onrcayci/goshell/internal/parser"
)

func Interpreter(argc int, argv []string) {
	if argc == 0 {
		return
	} else {
		switch argv[0] {
		case "help":
			help()
		case "quit":
			quit()
		case "set":
			err := set(argc, argv)
			if err != nil {
				fmt.Println(err.Error())
			}
		case "print":
			err := print(argc, argv)
			if err != nil {
				fmt.Println(err.Error())
			}
		case "run":
			err := run(argc, argv)
			if err != nil {
				fmt.Println(err.Error())
			}
		default:
			fmt.Printf("%s: command not found\n", argv[0])
		}
	}
}

func help() {
	helpText := `Go Shell v0.0.1
Available Commands:
COMMAND				DESCRIPTION

help				Displays all the commands
quit				Exits / terminates the shell with "Bye!"
set VAR STRING			Assigns a value to shell memory
print VAR			Displays the STRING assigned to VAR
run SCRIPT.TXT			Executes the file SCRIPT.TXT
`
	fmt.Println(helpText)
}

func quit() {
	fmt.Println("Bye!")
	os.Exit(0)
}

func set(argc int, args []string) error {
	if argc < 3 {
		return errors.New("missing arguments!\nusage: set VAR VALUE")
	}
	newVar := memory.NewMemoryItem(args[1], args[2])
	newVar.Set()
	return nil
}

func print(argc int, args []string) error {
	if argc < 2 {
		return errors.New("missing arguments!\nusage: print VAR")
	}
	varValue := memory.FindMemoryItem(args[1])
	if varValue == "" {
		return errors.New("variable does not exist")
	}
	fmt.Println(varValue)
	return nil
}

func run(argc int, args []string) error {
	if argc < 2 {
		return errors.New("missing arguments!\nusage: run SCRIPT.TXT")
	}
	script, err := os.Open(args[1])
	if err != nil {
		return err
	}
	reader := bufio.NewReader(script)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		argc, argv := parser.ParseInput(line)
		Interpreter(argc, argv)
	}
	return nil
}
