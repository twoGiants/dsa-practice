package generator

import (
	"fmt"
	"path/filepath"
	"slices"
)

type InputSource interface {
	Args() []string
}

type RequestImpl struct {
	command,
	pattern,
	difficulty string
	data DocsTemplateData
}

func NewRequest(
	source InputSource,
) (RequestImpl, error) {
	args := source.Args()
	if err := validate(args); err != nil {
		return RequestImpl{}, err
	}

	result := RequestImpl{
		args[0],
		args[1],
		args[2],
		NewDocsTemplateData(args[3]),
	}
	return result, nil
}

func validate(args []string) error {
	doValidate := func(attr string, member string, list []string) error {
		if !slices.Contains(list, member) {
			return fmt.Errorf("invalid %s: %s, allowed: %v", attr, member, list)
		}
		return nil
	}

	if len(args) != 4 {
		return fmt.Errorf("missing command, pattern, difficulty and title")
	}

	if err := doValidate(
		"command",
		args[0],
		[]string{"create", "delete"},
	); err != nil {
		return err
	}

	if err := doValidate(
		"pattern",
		args[1],
		[]string{"arrays", "bit-manipulation", "stack"},
	); err != nil {
		return err
	}

	if err := doValidate(
		"difficulty",
		args[2],
		[]string{"easy", "medium", "hard"},
	); err != nil {
		return err
	}

	return nil
}

func (r RequestImpl) Create() bool {
	return r.command == "create"
}

func (r RequestImpl) TemplateData() DocsTemplateData {
	return r.data
}

func (r RequestImpl) TargetPath(root string) string {
	return filepath.Join(
		root,
		r.pattern,
		r.difficulty,
		r.data.SmallTitle,
	)
}
