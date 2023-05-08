package utils

import (
	"fmt"
	"os"
	"time"
)

var file *os.File

type CustomLogger struct {
}

func (i *CustomLogger) Write(p []byte) (int, error) {
	var err error
	file, err = i.GetFile()
	if err != nil {
		return 0, err
	}
	return file.Write(p)
}

func (i *CustomLogger) GetFile() (*os.File, error) {
	currentDate := time.Now()
	currentFilename := fmt.Sprintf("logs/%s", currentDate.Format("2006-01-02.txt"))
	if file == nil {
		return os.OpenFile(currentFilename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	}
	lastFilename := file.Name()
	if currentFilename != lastFilename {
		err := file.Close()
		if err != nil {
			return nil, err
		}
		return os.OpenFile(currentFilename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	}
	return file, nil
}
