package boilerplate

import (
	"fmt"
	"path/filepath"
	"strings"
	"text/template"
)

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

func Generate(conf Config) error {
	if conf.create() {
		if err := DocsBoilerplateV2(conf); err != nil {
			return err
		}
	}

	if conf.delete() {
		if err := DeleteExercise(
			conf.targetDir,
			conf.exercise.pattern,
			conf.exercise.difficulty,
			conf.exercise.title,
		); err != nil {
			return err
		}
	}

	fmt.Println("Done.")

	return nil
}

func DocsBoilerplateV2(conf Config) error {
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

func DocsBoilerplate(exerciseRoot, tmplPath, pattern, difficulty, title string) error {
	tmpl, err := LoadDocs(tmplPath)
	if err != nil {
		return err
	}

	docs, err := Docs(tmpl, title)
	if err != nil {
		return err
	}

	exercisePath := filepath.Join(exerciseRoot, pattern, difficulty, smallKebab(title))
	if err := CreateExerciseDirectories(exercisePath); err != nil {
		return err
	}

	docsFileName := smallKebab(title) + ".md"
	boilerplatePath := filepath.Join(exercisePath, docsFileName)
	if err := StoreDocs(docs, boilerplatePath); err != nil {
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
