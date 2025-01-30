package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	fmt.Println("--- DSA Practice ---")
	processArgs()
}

func processArgs() {
	args := os.Args[1:]
	if len(args) != 3 {
		printUsage()
		panic(errors.New("missing pattern, difficulty and title"))
	}

	pattern := args[0]
	fmt.Println("pattern:", pattern)

	difficulty := args[1]
	fmt.Println("difficulty:", difficulty)

	title := args[2]
	fmt.Println("title:", title)
}

func printUsage() {
	fmt.Println("Usage: go run . <pattern> <difficulty> <title>")
}
