package logger

import (
	"errors"
	"fmt"
	"os"
	"path"
)

func getFilePointer(filePath string) (*os.File, error) {
	fmt.Printf("logger.getFilePointer: Getting pointer for file: %s\n", filePath)

	info, err := os.Stat(filePath)
	if err == nil {
		if info.IsDir() {
			return nil, errors.New("file path is a directory")
		}
		fmt.Printf("logger.getFilePointer: File already present.\n")
		return os.OpenFile(filePath, os.O_WRONLY, os.ModeAppend)
	}

	if !os.IsNotExist(err) {
		return nil, err
	}

	fmt.Printf("logger.getFilePointer: File absent. Creating...\n")
	err = os.MkdirAll(path.Dir(filePath), os.ModePerm)
	if err != nil && !os.IsExist(err) {
		return nil, fmt.Errorf("error while creating file: %s", err.Error())
	}

	return os.Create(filePath)
}
