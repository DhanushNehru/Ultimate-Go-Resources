package utils

import (
	"fmt"
	"time"
)

const (
	ColorReset  = "\033[0m"
	ColorRed    = "\033[31m"
	ColorGreen  = "\033[32m"
	ColorYellow = "\033[33m"
	ColorBlue   = "\033[34m"
)

func PrintSuccess(message string) {
	fmt.Println(string(ColorGreen), message, string(ColorReset))
}

func PrintError(message string) {
	fmt.Println(string(ColorRed), message, string(ColorReset))
}

func PrintInfo(message string) {
	fmt.Println(string(ColorYellow), message, string(ColorReset))
}

func ValidateDate(dateStr string) bool {
	if dateStr == "" {
		return false
	}
	_, err := time.Parse("2006-01-02", dateStr)
	return err == nil
}
