package tasks

import (
	"time"
)

type Task struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Status   string `json:"status"`   
	Due      string `json:"due"`      
	Priority string `json:"priority"` 
}

func NewTask(title string, due string, priority string) Task {
	if priority != "low" && priority != "medium" && priority != "high" {
		priority = "medium"
	}

	if due != "" {
		_, err := time.Parse("2006-01-02", due)
		if err != nil {
			due = ""
		}
	}

	return Task{
		Title:    title,
		Status:   "pending",
		Due:      due,
		Priority: priority,
	}
}
