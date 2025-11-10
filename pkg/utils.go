package deepspec

import (
	"fmt"
	"os"
	"time"
)

const logFilePath = "out.log"

// LogToFile writes a log message to a file for debugging
func LogToFile(format string, args ...interface{}) {
	f, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return
	}
	defer f.Close()

	timestamp := time.Now().Format("2006-01-02 15:04:05")
	message := fmt.Sprintf(format, args...)
	logLine := fmt.Sprintf("[%s] %s\n", timestamp, message)
	f.WriteString(logLine)
}
