package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	api "github.com/marcanodev/pokedex/internal/pokeapi"
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
	mapLoc
}

type mapLoc struct {
	Next string
	Prev interface{}
}

var validCmd = make(map[string]cmd)

var URL = "https://pokeapi.co/api/v2/location/"

var mapNextCmd cmd
var mapPrevCmd cmd

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
		case "map":
			locs := &api.Locations{}
			data, err := api.GetJSON(URL)

			if err != nil {
				log.Fatal(err)
			}

			err = json.Unmarshal(data, locs)
			if err != nil {
				log.Fatal(err)
			}

			mapNextCmd.Next = *locs.Next
			mapPrevCmd.Prev = locs.Prev
			for _, l := range locs.Results {
				fmt.Fprintln(os.Stdout, string(*l.Name))
			}

			URL = mapNextCmd.Next
		default:
			fmt.Fprint(os.Stdout, usageMsg)
		}

	}
}

func init() {
	help := addCmd("help", helpDesc, helpCB)
	exit := addCmd("exit", exitDesc, exitCB)
	mapNextCmd = addCmd("map",
		"Shows next 20 locations areas of the Pokemon world", mapCB)
	mapPrevCmd = addCmd("mapb",
		"Shows the prev 20 locations of the Pokemon world", mapCB)

	usageMsg = fmt.Sprintf(`Welcome to the %s!
Usage:

%s: %s
%s: %s
%s: %s

`, progName,
		help.name, help.description,
		exit.name, exit.description,
		mapNextCmd.name, mapNextCmd.description)
}

func addCmd(name string, desc string, cb func() error) cmd {
	name = strings.ToLower(name)
	validCmd[name] = cmd{
		name:        name,
		description: desc,
		callback:    cb,
	}

	return validCmd[name]
}

func helpCB() error {
	return nil
}

func exitCB() error {
	return nil
}

func mapCB() error {
	return nil
}
