package main

import (
	"dsa/utils/generator"
	"errors"
	"fmt"
	"os"
	"slices"
)

func main() {
	fmt.Println("--- DSA Practice ---")
	info, err := processArgs()
	if err != nil {
		panic(err)
	}

	tempDir := "temp"
	if _, err := os.Stat(tempDir); err != nil {
		panic(fmt.Errorf("%v, create %s and try again", err, tempDir))
	}

	if info.command == "create" {
		tmplPath := "utils/generator/docs.gotmpl"
		if err := generator.DocsBoilerplate(tempDir, tmplPath, info.pattern, info.difficulty, info.title); err != nil {
			panic(err)
		}
	}

	if info.command == "delete" {
		if err := generator.DeleteExercise(tempDir, info.pattern, info.difficulty, info.title); err != nil {
			panic(err)
		}
	}
}

type exercise struct {
	command, pattern, difficulty, title string
}

func processArgs() (exercise, error) {
	var anExercise exercise
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

	return exercise{
		command,
		pattern,
		difficulty,
		title,
	}, nil
}

func printUsage() {
	fmt.Println("Usage: go run . <command> <pattern> <difficulty> <title>")
}
