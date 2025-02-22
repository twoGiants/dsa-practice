package boilerplate

import (
	"errors"
	"fmt"
	"os"
	"slices"
)

type metadata struct {
	command, pattern, difficulty, title string
}

func ProcessArgs() (metadata, error) {
	var anExercise metadata
	args := os.Args[1:]
	if len(args) != 4 {
		printUsage()
		return anExercise, errors.New("missing command, pattern, difficulty and title")
	}

	commands := []string{"create", "delete"}
	command := args[0]
	if !slices.Contains(commands, command) {
		return anExercise, fmt.Errorf("invalid command: %s, allowed: %v", command, commands)
	}
	fmt.Println("command:", command)

	patterns := []string{"arrays", "bit-manipulation", "stack"}
	pattern := args[1]
	if !slices.Contains(patterns, pattern) {
		return anExercise, fmt.Errorf("invalid pattern: %s, allowed: %v", pattern, patterns)
	}
	fmt.Println("pattern:", pattern)

	difficulties := []string{"easy", "medium", "hard"}
	difficulty := args[2]
	if !slices.Contains(difficulties, difficulty) {
		return anExercise, fmt.Errorf("invalid difficulty: %s, allowed: %v", difficulty, difficulties)
	}
	fmt.Println("difficulty:", difficulty)

	title := args[3]
	fmt.Println("title:", title)

	return metadata{
		command,
		pattern,
		difficulty,
		title,
	}, nil
}

func printUsage() {
	fmt.Println("Usage: go run . <command> <pattern> <difficulty> <title>")
}
