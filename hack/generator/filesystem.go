package generator

import (
	"fmt"
	"os"
)

type Filesystem interface {
	Load(from string) (string, error)
	Save(data, to string) error
}

type FilesystemImpl struct{}

func NewFilesystem() FilesystemImpl {
	return FilesystemImpl{}
}

func (fsi FilesystemImpl) Load(from string) (string, error) {
	result, err := os.ReadFile(from)
	if err != nil {
		return "", err
	}

	return string(result), err
}

func (fsi FilesystemImpl) Save(data, to string) error {
	fmt.Println(data)
	fmt.Println(to)
	if err := os.WriteFile(to, []byte(data), 0600); err != nil {
		return err
	}

	return nil
}
