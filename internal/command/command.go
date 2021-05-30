package command

import (
	"errors"
	"fmt"
	"os"

	"github.com/onrcayci/goshell/internal/memory"
)

func Help() {
	helpText := `Go Shell v0.0.1
Available Commands:
COMMAND				DESCRIPTION

help				Displays all the commands
quit				Exits / terminates the shell with "Bye!"
set VAR STRING		Assigns a value to shell memory
print VAR			Displays the STRING assigned to VAR
run SCRIPT.TXT		Executes the file SCRIPT.TXT
`
	fmt.Println(helpText)
}

func Quit() {
	fmt.Println("Bye!")
	os.Exit(0)
}

func Set(argc int, args []string) error {
	if argc < 3 {
		return errors.New("missing arguments!\nusage: set VAR VALUE")
	}
	newVar := memory.NewMemoryItem(args[1], args[2])
	newVar.Set()
	return nil
}
