package command

import (
	"fmt"
	"os"
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

func Exit() {
	os.Exit(0)
}
