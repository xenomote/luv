package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
)

const stdin = "-"

var params = regexp.MustCompile(`\S+`)

func main() {
	args := os.Args[1:] // drop program name

	if len(args) == 0 {
		usage("no arguments")
	}

	cmd := "main"
	_, matches := map[string]any{
		"select":   nil,
		"extract":  nil,
		"parse":    nil,
		"validate": nil,
	}[args[0]]

	if matches {
		cmd = args[0] 	// use subcommand
		args = args[1:] // drop subcommand name
	}

	if len(args) == 0 {
		usage("no arguments to", cmd)
	}

	// TODO: use scanner so full input is not buffered
	if len(args) == 1 && args[0] == stdin {
		input, err := bufio.NewReader(os.Stdin).ReadString(0)
		if err != nil && err != io.EOF {
			usage("failed to read stdin:", err)
		}
		args = params.FindAllString(input, -1)
	}

	for _, a := range args {
		if a == "-" {
			usage("cannot use", stdin, "as a target")
		}
	}
	
	succeed("run", cmd, "on", args)
}

func succeed(msg ...any) {
	fmt.Println(msg...)
	os.Exit(0)
}

func usage(msg ...any) {
	fmt.Println(msg...)
	os.Exit(1)
}
