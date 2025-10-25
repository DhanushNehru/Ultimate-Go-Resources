package tasks

import (
	"encoding/json"
	"errors"
	"os"
)

var taskFile = "tasks.json"

func LoadTasks() ([]Task, error) {
	if _, err := os.Stat(taskFile); errors.Is(err, os.ErrNotExist) {
		// File doesn't exist yet, return empty slice
		return []Task{}, nil
	}

	data, err := os.ReadFile(taskFile)
	if err != nil {
		return nil, err
	}

	var tasks []Task
	err = json.Unmarshal(data, &tasks)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func SaveTasks(tasks []Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile(taskFile, data, 0644)
	if err != nil {
		return err
	}

	return nil
}
