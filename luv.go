package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:] // drop program name

	if len(args) == 0 {
		usage("no arguments")
	}

	sub := args[0]
	_, matches := map[string]any{
		"select":   nil,
		"extract":  nil,
		"parse":    nil,
		"validate": nil,
	}[sub]
	if !matches {
		succeed("run main process on:", args)
	}

	args = args[1:] // drop subcommand name
	succeed("run", sub, "on", args)
}

func succeed(msg ...any) {
	fmt.Println(msg...)
	os.Exit(0)
}

func usage(msg ...any) {
	fmt.Println(msg...)
	os.Exit(1)
}
