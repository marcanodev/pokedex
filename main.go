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

func (m *mapLoc) Request(url *string) (l *api.Payload, e error) {
	l = &api.Payload{}

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
var _map *cmd
var mapB *cmd
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
			resp, err := _map.Request(&URL)
			if err != nil && resp.Count <= 0 {
				log.Fatal(err)
			}

			_map.Next = resp.Next
			mapB.Prev = resp.Prev
			for _, l := range resp.Results {
				fmt.Fprintln(os.Stdout, string(*l.Name))
			}

			URL = *_map.Next
		case "mapb":
			resp, err := mapB.Request(mapB.Prev)
			if err != nil {
				log.Fatal(err)
			}

			mapB.Prev = resp.Prev
			for _, l := range resp.Results {
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
	_map = addCmd("map",
		"Shows next 20 locations areas of the Pokemon world", mapCB)
	mapB = addCmd("mapb",
		"Shows the prev 20 locations of the Pokemon world", mapCB)

	usageMsg = fmt.Sprintf(`Welcome to the %s!
Usage:

%s: %s
%s: %s
%s: %s
%s: %s

`, progName,
		helpCmd.Name, helpCmd.Description,
		exitCmd.Name, exitCmd.Description,
		_map.Name, _map.Description,
		mapB.Name, mapB.Description)
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
