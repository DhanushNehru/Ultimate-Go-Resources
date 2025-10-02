package tasks

import (
	"errors"
)

func AddTask(task Task) error {
	tasks, err := LoadTasks()
	if err != nil {
		return err
	}

	// Assign ID
	if len(tasks) == 0 {
		task.ID = 1
	} else {
		task.ID = tasks[len(tasks)-1].ID + 1
	}

	tasks = append(tasks, task)
	return SaveTasks(tasks)
}

func ListTasks(showAll bool) ([]Task, error) {
	tasks, err := LoadTasks()
	if err != nil {
		return nil, err
	}

	if showAll {
		return tasks, nil
	}

	var pending []Task
	for _, t := range tasks {
		if t.Status != "completed" {
			pending = append(pending, t)
		}
	}
	return pending, nil
}

func CompleteTask(id int) error {
	tasks, err := LoadTasks()
	if err != nil {
		return err
	}

	found := false
	for i, t := range tasks {
		if t.ID == id {
			tasks[i].Status = "completed"
			found = true
			break
		}
	}

	if !found {
		return errors.New("task not found")
	}

	return SaveTasks(tasks)
}

func RemoveTask(id int) error {
	tasks, err := LoadTasks()
	if err != nil {
		return err
	}

	index := -1
	for i, t := range tasks {
		if t.ID == id {
			index = i
			break
		}
	}

	if index == -1 {
		return errors.New("task not found")
	}

	tasks = append(tasks[:index], tasks[index+1:]...)
	return SaveTasks(tasks)
}
