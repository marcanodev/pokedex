package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	progName = "Pokedex"
)

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
			fmt.Println("help message")
		}
	}
}
