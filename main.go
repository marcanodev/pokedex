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
	Name        string
	Description string
	Callback    func() error
	mapLoc
}

type mapLoc struct {
	Next *string
	Prev *string
}

func (m *mapLoc) Request(url *string) (l *api.Locations, e error) {
	l = &api.Locations{}

	if url != nil {
		data, err := api.GetJSON(*url)

		if err != nil {
			e = err
		}

		err = json.Unmarshal(data, l)
		if err != nil {
			e = err
		}
	}

	return
}

var validCmd = make(map[string]cmd)
var mapNextCmd *cmd
var mapPrevCmd *cmd
var helpCmd *cmd
var exitCmd *cmd

var URL = "https://pokeapi.co/api/v2/location/"

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
			locs, err := mapNextCmd.Request(&URL)
			if err != nil && locs.Count <= 0 {
				log.Fatal(err)
			}

			mapNextCmd.Next = locs.Next
			mapPrevCmd.Prev = locs.Prev
			for _, l := range locs.Results {
				fmt.Fprintln(os.Stdout, string(*l.Name))
			}

			URL = *mapNextCmd.Next
		case "mapb":
			locs, err := mapPrevCmd.Request(mapPrevCmd.Prev)
			if err != nil {
				log.Fatal(err)
			}

			mapPrevCmd.Prev = locs.Prev
			for _, l := range locs.Results {
				fmt.Fprintln(os.Stdout, string(*l.Name))
			}
		default:
			fmt.Fprint(os.Stdout, usageMsg)
		}

	}
}

func init() {
	helpCmd = addCmd("help", helpDesc, helpCB)
	exitCmd = addCmd("exit", exitDesc, exitCB)
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
		helpCmd.Name, helpCmd.Description,
		exitCmd.Name, exitCmd.Description,
		mapNextCmd.Name, mapNextCmd.Description)
}

func addCmd(name string, desc string, cb func() error) *cmd {
	c := &cmd{
		Name:        strings.ToLower(name),
		Description: desc,
		Callback:    cb,
	}
	validCmd[name] = *c

	return c
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
