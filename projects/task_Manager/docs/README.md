# Go Task Manager CLI 

A simple, open-source **command-line task manager** written in Go.  
Manage your tasks, mark them as completed, and track priorities – all from your terminal!

---

## Features

- Add tasks with **title, due date, and priority**  
- List pending tasks or all tasks  
- Mark tasks as completed  
- Remove tasks  
- JSON-based persistent storage   

---

## Folder Purpose Explained

cmd/taskmanager/main.go → The CLI entry point; parses commands (add, remove, list, complete).
pkg/tasks/ → Core logic for tasks and storage.
task.go → Defines Task struct (title, description, status, due date, priority).
manager.go → CRUD operations: add, remove, list, complete, search tasks.
storage.go → Save/load tasks to a JSON file (later can be extended to SQLite).
pkg/utils/helpers.go → Small helpers: string validation, date parsing, color formatting for CLI output.
internal/config/config.go → For app configuration (default storage file, date format, etc.).
docs/ → Project documentation and Hacktoberfest instructions.
examples/ → Predefined tasks for testing or demos.