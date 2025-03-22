package boilerplate

import (
	"fmt"
	"path/filepath"
	"slices"
	"strings"
	"text/template"
)

type metadata struct {
	command,
	pattern,
	difficulty,
	title string
}

func NewMetadata(
	command,
	pattern,
	difficulty,
	title string,
) metadata {
	return metadata{
		command,
		pattern,
		difficulty,
		title,
	}
}

func ValidMetadata(m metadata) (metadata, error) {
	if err := validate(
		"command",
		m.command,
		[]string{"create", "delete"},
	); err != nil {
		return m, err
	}

	if err := validate(
		"pattern",
		m.pattern,
		[]string{"arrays", "bit-manipulation", "stack"},
	); err != nil {
		return m, err
	}

	if err := validate(
		"difficulty",
		m.difficulty,
		[]string{"easy", "medium", "hard"},
	); err != nil {
		return m, err
	}

	return m, nil
}

func validate(attr string, member string, list []string) error {
	if !slices.Contains(list, member) {
		return fmt.Errorf("invalid %s: %s, allowed: %v", attr, member, list)
	}
	fmt.Printf("%s: %s\n", attr, member)

	return nil
}

type Config struct {
	docsTmplPath string
	targetDir    string
	exercise     metadata
}

func NewConfig(
	docsTmplPath,
	targetDir string,
	exercise metadata,
) Config {
	return Config{
		docsTmplPath,
		targetDir,
		exercise,
	}
}

func (c *Config) create() bool {
	return c.exercise.command == "create"
}

func (c *Config) delete() bool {
	return c.exercise.command == "delete"
}

func (c *Config) exercisePath() string {
	return filepath.Join(c.targetDir, c.exercise.pattern, c.exercise.difficulty, smallKebab(c.exercise.title))
}

func (c *Config) docsFilePath() string {
	docsFileName := smallKebab(c.exercise.title) + ".md"
	return filepath.Join(c.exercisePath(), docsFileName)
}

type Generator struct {
	Config
}

func NewGenerator(c Config) *Generator {
	return &Generator{c}
}

func (g *Generator) Run() error {
	if g.create() {
		if err := DocsBoilerplate(g.Config); err != nil {
			return err
		}
	}

	if g.delete() {
		if err := DeleteExercise(
			g.targetDir,
			g.exercise.pattern,
			g.exercise.difficulty,
			g.exercise.title,
		); err != nil {
			return err
		}
	}

	fmt.Println("Done.")

	return nil
}

func DocsBoilerplate(conf Config) error {
	tmpl, err := LoadDocs(conf.docsTmplPath)
	if err != nil {
		return err
	}

	docs, err := Docs(tmpl, conf.exercise.title)
	if err != nil {
		return err
	}

	if err := CreateExerciseDirectories(conf.exercisePath()); err != nil {
		return err
	}

	if err := StoreDocs(docs, conf.docsFilePath()); err != nil {
		return err
	}

	return nil
}

func Docs(notParsed, title string) (string, error) {
	tmpl, err := template.New("docs").Parse(notParsed)
	if err != nil {
		return "", err
	}

	small := smallKebab(title)

	data := struct {
		Title      string
		SmallTitle string
	}{
		Title:      title,
		SmallTitle: small,
	}

	var result strings.Builder
	if err := tmpl.Execute(&result, data); err != nil {
		return "", err
	}

	return result.String(), nil
}

func smallKebab(title string) string {
	small := strings.ToLower(title)

	if len(strings.Split(small, " ")) > 1 {
		small = strings.ReplaceAll(small, " ", "-")
	}
	return small
}
