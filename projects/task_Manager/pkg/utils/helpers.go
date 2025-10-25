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
	fmt.Printf("%s%s%s\n", ColorGreen, message, ColorReset)
}
func PrintError(message string) {
	fmt.Printf("%s%s%s\n", ColorRed, message, ColorReset)
}
func PrintInfo(message string) {
	fmt.Printf("%s%s%s\n", ColorYellow, message, ColorReset)
}
func ValidateDate(dateStr string) bool {
	if dateStr == "" {
		return false
	}
	_, err := time.Parse("2006-01-02", dateStr)
	return err == nil
}
