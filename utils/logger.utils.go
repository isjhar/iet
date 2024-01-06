package utils

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

var file *os.File

const LogInfoLevel = 1
const LogWarningLevel = 2
const LogErrorLevel = 3

var LogLevel = LogInfoLevel

type CustomLogger struct {
}

func (i *CustomLogger) Write(p []byte) (int, error) {
	var err error
	file, err = i.GetFile()
	if err != nil {
		return 0, err
	}
	logMultiwriter := io.MultiWriter(os.Stdout, file)
	return logMultiwriter.Write(p)
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

func LogInfo(format string, v ...any) {
	if LogLevel > LogInfoLevel {
		return
	}
	log.SetPrefix("INFO ")
	log.Printf(format, v...)
}

func LogWarning(format string, v ...any) {
	if LogLevel > LogWarningLevel {
		return
	}
	log.SetPrefix("INFO ")
	log.Printf(format, v...)
}

func LogError(format string, v ...any) {
	if LogLevel > LogErrorLevel {
		return
	}
	log.SetPrefix("INFO ")
	log.Printf(format, v...)
}
