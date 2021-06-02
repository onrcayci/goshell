package command

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/onrcayci/goshell/internal/cpu"
	"github.com/onrcayci/goshell/internal/kernel"
	"github.com/onrcayci/goshell/internal/memory"
	"github.com/onrcayci/goshell/internal/parser"
	"github.com/onrcayci/goshell/internal/pcb"
	"github.com/onrcayci/goshell/internal/ram"
)

var exitFlag bool = true

func Execute(argc int, argv []string) {
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
	case "exec":
		err := exec(argc, argv)
		if err != nil {
			fmt.Println(err.Error())
		}
	default:
		fmt.Printf("%s: command not found\n", argv[0])
	}
}

// Function help which implements the "help" shell command.
// Returns the help text which outputs all of the supported commands
// and their description.
func help() {
	helpText := `Go Shell v1.0
Available Commands:
COMMAND				DESCRIPTION

help				Displays all the commands
quit				Exits / terminates the shell with "Bye!"
set VAR STRING			Assigns a value to shell memory
print VAR			Displays the STRING assigned to VAR
run SCRIPT.TXT			Executes the file SCRIPT.TXT
exec p1 p2 p3			Executes concurrent programs: >> exec prog.txt prog2.txt
`
	fmt.Println(helpText)
}

// Function quit which implements the "quit" shell command.
// Prints out the message "Bye!" on the screen and end the appliction runtime
// using the function os.Exit(0).
func quit() {
	fmt.Println("Bye!")
	if exitFlag {
		os.Exit(0)
	}
}

// Function set which implements the "set" shell command.
// Creates and saves a new shell environment variable into the runtime if the
// runtime does not exist. If the variable exists, it updates the value of it
// using the new value passed into the command.
// Returns an error if the number of arguments is less than 3 (i.e. "set VAR VALUE").
func set(argc int, args []string) error {
	if argc < 3 {
		return errors.New("missing arguments!\nusage: set VAR VALUE")
	}
	memory.NewMemoryItem(args[1], args[2])
	return nil
}

// Function print which implements the "print" shell command.
// Displays the value of a saved shell environment variable if it exists.
// Otherwise, displays the error message "variable does not exist".
// Returns an error if the number of arguments is less than 2 (i.e. "print VAR")
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

// Function run which implements the "run" shell command.
// This function enables users to run shell scripts with supported commands.
// Returns an error if the number of arguments is less than 2 (i.e. "run SCRIPT.TXT").
// BUG(onrcayci): The function parser.ParseInput parses the filename into 3 tokens, i.e., [<filename> "." <file extension>].
func run(argc int, args []string) error {
	if argc < 2 {
		return errors.New("missing arguments!\nusage: run SCRIPT.TXT")
	}
	// due to the file name tokenization bug, the filename is provided using 3 arguments from args:
	// args[1] = file name, args[2] = ".", args[3] = file extension.
	script, err := os.Open(args[1] + args[2] + args[3])
	if err != nil {
		return err
	}
	exitFlag = false
	reader := bufio.NewReader(script)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			exitFlag = true
			return err
		}
		argc, argv := parser.ParseInput(line)
		Execute(argc, argv)
	}
	exitFlag = true
	return nil
}

func runForQuanta(c *cpu.CPU) {
	for c.Quanta != 0 {
		c.IR = ram.RAM[c.IP]
		c.IP++
		c.Quanta--
		argc, argv := parser.ParseInput(c.IR)
		exitFlag = false
		Execute(argc, argv)
		exitFlag = true
	}
}

func scheduler() {
	for kernel.ReadyQueue.Front() != nil {
		if kernel.RuntimeCPU.Quanta == 0 || kernel.RuntimeCPU.IR == "" {
			currentPCBElement := kernel.ReadyQueue.Front()
			currentPCB := currentPCBElement.Value.(*pcb.PCB)
			kernel.RuntimeCPU.IP = currentPCB.PC
			quanta := (currentPCB.End - 1) - currentPCB.PC
			if quanta == 0 {
				ram.FreeRAM(currentPCB.Start, currentPCB.End)
				kernel.ReadyQueue.Remove(currentPCBElement)
				continue
			} else if quanta > 2 {
				quanta = 2
			}
			kernel.RuntimeCPU.Quanta = quanta
			runForQuanta(kernel.RuntimeCPU)
			currentPCB.PC = kernel.RuntimeCPU.IP
			kernel.ReadyQueue.Remove(currentPCBElement)
			kernel.ReadyQueue.PushBack(currentPCB)
		}
	}
}

func exec(argc int, argv []string) error {
	scripts := argv[1:]
	if len(scripts) == 0 {
		return errors.New("missing arguments!\nusage: exec p1.txt [p2.txt] [p3.txt]")
	}
	for i := 0; (3 * i) < len(scripts); i++ {
		err := kernel.MyInit(scripts[3*i] + scripts[3*i+1] + scripts[3*i+2])
		if err != nil {
			ram.FreeRAM(0, 1000)
			return err
		}
	}
	exitFlag = false
	scheduler()
	exitFlag = true
	return nil
}
