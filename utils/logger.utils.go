package utils

import (
	"fmt"
	"io"
	"log"
	"os"
	"runtime"
	"strings"
	"time"
)

var file *os.File

const LogInfoLevel = "INFO"
const LogWarningLevel = "WARNING"
const LogErrorLevel = "ERROR"

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
	if LogLevel == LogWarningLevel || LogLevel == LogErrorLevel {
		return
	}
	logMessagef("INFO", format, v...)
}

func LogInfoln(v ...any) {
	if LogLevel == LogWarningLevel || LogLevel == LogErrorLevel {
		return
	}
	logMessageln("INFO", v...)
}

func LogWarning(format string, v ...any) {
	if LogLevel == LogErrorLevel {
		return
	}
	logMessagef("WARNING", format, v...)
}

func LogWarningln(v ...any) {
	if LogLevel == LogErrorLevel {
		return
	}
	logMessageln("WARNING", v...)
}

func LogError(format string, v ...any) {
	logMessagef("ERROR", format, v...)
}

func logMessagef(level string, format string, v ...any) {
	message := fmt.Sprintf(format, v...)
	if strings.Contains(message, "context canceled") {
		return
	}
	// Get the file and line number
	_, file, line, ok := runtime.Caller(2) // Caller(2) to get the function that called logMessage
	if !ok {
		file = "???"
		line = 0
	}
	message = fmt.Sprintf("%s:%d %s", file, line, message)

	log.SetPrefix(level + " ")
	log.Println(message)
}

func logMessageln(level string, v ...any) {
	message := fmt.Sprint(v...)
	if strings.Contains(message, "context canceled") {
		return
	}
	// Get the file and line number
	_, file, line, ok := runtime.Caller(2) // Caller(2) to get the function that called logMessage
	if !ok {
		file = "???"
		line = 0
	}

	message = fmt.Sprintf("%s:%d %s", file, line, message)
	log.SetPrefix(level + " ")
	log.Println(message)
}

func GetLogLevel() string {
	return GetEnvironmentVariable("LOG_LEVEL", LogLevel)
}
