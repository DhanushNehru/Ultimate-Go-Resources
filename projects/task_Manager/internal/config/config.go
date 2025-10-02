package config

import (
	"os"
)

type Config struct {
	TaskFile   string
	DateFormat string
}

func LoadConfig() Config {
	taskFile := "tasks.json"

	if envFile := os.Getenv("TASK_FILE"); envFile != "" {
		taskFile = envFile
	}

	return Config{
		TaskFile:   taskFile,
		DateFormat: "2006-01-02",
	}
}
