package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	progName = "Pokedex"
	helpDesc = "Displays a help message"
	exitDesc = "Exit the pokedex"
)

type cmd struct {
	name        string
	description string
	callback    func() error
}

var validCmd = make(map[string]cmd)

var input string
var scanner *bufio.Scanner
var prompt string = fmt.Sprintf("%s > ", strings.ToLower(progName))
var usageMsg string

func main() {
	scanner = bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(prompt)
		scanner.Scan()
		input = strings.ToLower(scanner.Text())

		switch input {
		case "exit":
			os.Exit(0)
		case "help":
			fmt.Fprint(os.Stdout, usageMsg)
		default:
			fmt.Fprint(os.Stdout, usageMsg)
		}

	}
}

func init() {
	validCmd["help"] = cmd{
		name:        "help",
		description: helpDesc,
		callback:    helpCmd,
	}

	validCmd["exit"] = cmd{
		name:        "exit",
		description: exitDesc,
		callback:    exitCmd,
	}

	usageMsg = fmt.Sprintf(`Welcome to the %s!
Usage:

%s: %s
%s: %s

`, progName,
		validCmd["help"].name, validCmd["help"].description,
		validCmd["exit"].name, validCmd["exit"].description)
}

func helpCmd() error {
	return nil
}

func exitCmd() error {
	return nil
}
