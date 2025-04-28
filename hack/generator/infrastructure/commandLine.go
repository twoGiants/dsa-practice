package infrastructure

import "os"

type CommandLine struct{}

func (c CommandLine) Args() []string {
	return os.Args[1:]
}
