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

	exercise, err := boilerplate.ProcessArgs()
	if err != nil {
		panic(err)
	}

	conf := boilerplate.NewConfig(
		"utils/boilerplate/docs.gotmpl",
		tempDir,
		exercise,
	)

	if err := boilerplate.Generate(conf); err != nil {
		panic(err)
	}
}
