package logger

import (
	"fmt"
	"time"
)

const (
	ColorReset  = "\033[0m"
	ColorRed    = "\033[31m"
	ColorYellow = "\033[33m"
	ColorGreen  = "\033[32m"
	ColorBlue   = "\033[34m"
)

func Info(msg string, args ...interface{}) {
	log("INFO", ColorGreen, msg, args...)
}

func Warn(msg string, args ...interface{}) {
	log("WARN", ColorYellow, msg, args...)
}

func Error(msg string, args ...interface{}) {
	log("ERROR", ColorRed, msg, args...)
}

func log(level, color, msg string, args ...interface{}) {
	timestamp := time.Now().Format(time.RFC3339)
	formatted := fmt.Sprintf(msg, args...)
	fmt.Printf("%s[%s] %s | %s%s\n", color, level, timestamp, formatted, ColorReset)
}
