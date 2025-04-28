package generator

import (
	"os"
	"path/filepath"
)

type FilesystemImpl struct{}

func NewFilesystem() FilesystemImpl {
	return FilesystemImpl{}
}

func (fsi FilesystemImpl) Load(from string) (string, error) {
	result, err := os.ReadFile(filepath.Clean(from))
	if err != nil {
		return "", err
	}

	return string(result), err
}

func (fsi FilesystemImpl) Save(data, to string) error {
	if err := os.MkdirAll(filepath.Dir(to), 0750); err != nil {
		return err
	}

	return os.WriteFile(to, []byte(data), 0600)
}

func (fsi FilesystemImpl) Delete(from string) error {
	return os.RemoveAll(from)
}
