package store

import (
	"os"
)

var dir = "internal/data/"

func EnsureFile(fileDir string) (string, error) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.Mkdir(dir, 0755); err != nil {
			return "", err
		}
	}
	filePath := dir + fileDir
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		if err := os.WriteFile(filePath, []byte("{}"), 0644); os.IsNotExist(err) {
			return "", err
		}
	}
	return filePath, nil
}
