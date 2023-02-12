package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	progName = "Pokedex"
	helpMsg  = "Displays a help message"
	exitMsg  = "Exit the pokedex"
)

type cmd struct {
	name        string
	description string
	callback    func() error
}

var validCmd = make(map[string]cmd)

var input string
var scanner *bufio.Scanner
var prompt = fmt.Sprintf("%s > ", strings.ToLower(progName))

func main() {
	scanner = bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(prompt)
		scanner.Scan()
		input = scanner.Text()

		switch input {
		case "q":
			fmt.Println("Come back again soon!")
			os.Exit(0)
		case "h":
			fmt.Println(validCmd["help"].description)
		}
	}
}

func init() {
	validCmd["help"] = cmd{
		name:        "help",
		description: helpMsg,
		callback:    helpCmd,
	}

	validCmd["exit"] = cmd{
		name:        "exit",
		description: exitMsg,
		callback:    exitCmd,
	}
}

func helpCmd() error {
	return nil
}

func exitCmd() error {
	return nil
}
