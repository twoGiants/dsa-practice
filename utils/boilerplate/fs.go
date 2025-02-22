package boilerplate

import (
	"os"
	"path/filepath"
)

func LoadDocs(tmplPath string) (string, error) {
	result, err := os.ReadFile(tmplPath)
	if err != nil {
		return "", err
	}

	return string(result), err
}

func StoreDocs(data, path string) error {
	if err := os.WriteFile(path, []byte(data), 0644); err != nil {
		return err
	}
	return nil
}

func DeleteExercise(exerciseRoot, pattern, difficulty, title string) error {
	path := filepath.Join(exerciseRoot, pattern, difficulty, smallKebab(title))
	if err := os.RemoveAll(path); err != nil {
		return err
	}

	return nil
}

func CreateExerciseDirectories(path string) error {
	if err := os.MkdirAll(path, 0755); err != nil {
		return err
	}

	return nil
}
