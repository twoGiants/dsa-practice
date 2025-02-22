package main

import (
	"dsa/utils/boilerplate"
	"fmt"
	"os"
)

func main() {
	fmt.Println("--- DSA Practice ---")

	tempDir := "temp"
	if _, err := os.Stat(tempDir); err != nil {
		panic(fmt.Errorf("%v, create %s and try again", err, tempDir))
	}

	args := os.Args[1:]
	if len(args) != 4 {
		printUsage()
		panic("missing command, pattern, difficulty and title")
	}

	metadata, err := boilerplate.ValidatedMetadata(
		boilerplate.NewMetadata(
			args[0],
			args[1],
			args[2],
			args[3],
		),
	)
	if err != nil {
		panic(err)
	}

	generator := boilerplate.NewGenerator(
		boilerplate.NewConfig(
			"utils/boilerplate/docs.gotmpl",
			tempDir,
			metadata,
		),
	)

	if err := generator.Run(); err != nil {
		panic(err)
	}
}

func printUsage() {
	fmt.Println("Usage: go run . <command> <pattern> <difficulty> <title>")
}
