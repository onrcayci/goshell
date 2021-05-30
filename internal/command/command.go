package command

import (
	"fmt"
	"os"
)

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
