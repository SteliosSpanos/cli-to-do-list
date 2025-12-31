# CLI To-Do List

A command-line task management application written in Go that provides a simple and efficient way to manage your daily tasks through terminal commands.

## Overview

This application allows users to manage tasks through a command-line interface with persistent JSON-based storage. Tasks can be added, deleted, marked as complete, and displayed with their current status.

## Features

- Add new tasks with descriptions
- Delete tasks by ID
- Mark tasks as completed
- Display all tasks with their status
- Persistent storage using JSON format
- Automatic ID assignment for new tasks
- Error handling with clear user feedback

## Project Structure

```
to-do-list/
├── cmd/
│   └── main/
│       └── main.go          # Application entry point
├── commands/
│   ├── add.go              # Add task functionality
│   ├── delete.go           # Delete task functionality
│   ├── complete.go         # Complete task functionality
│   └── show.go             # Display tasks functionality
├── pkg/
│   └── task.go             # Core task data structures and operations
├── tasks.json              # Task storage file (auto-generated)
└── go.mod                  # Go module definition
```

## Installation

### Prerequisites

- Go 1.25.5 or higher

### Build from Source

```bash
git clone https://github.com/SteliosSpanos/cli-to-do-list.git
cd cli-to-do-list
go build -o main cmd/main/main.go
```

## Usage

### Basic Syntax

```bash
./main <command> [arguments]
```

### Available Commands

#### Add a Task

Add a new task to your to-do list.

```bash
./main add "Task description"
```

Example:
```bash
./main add "Buy groceries"
```

#### Delete a Task

Remove a task from your list by its ID.

```bash
./main delete <id>
```

Example:
```bash
./main delete 1
```

#### Complete a Task

Mark a task as completed by its ID.

```bash
./main complete <id>
```

Example:
```bash
./main complete 2
```

#### Show All Tasks

Display all tasks with their IDs and status.

```bash
./main show
```

Output format:
```
=======================================================
                        YOUR TASKS
=======================================================

[ID 1] Buy groceries  |Completed|
[ID 2] Write documentation  |Not Completed|
=======================================================
```

## Data Storage

Tasks are stored in a `notes.json` file in the application directory. The file is automatically created on first use and follows this structure:

```json
{
  "tasks": [
    {
      "id": 1,
      "description": "Task description",
      "status": "Not Completed"
    }
  ]
}
```

## Technical Details

### Core Components

#### Task Structure

The application uses two primary data structures:

- **Task**: Represents a single task with ID, description, and status
- **TaskList**: Contains a collection of tasks

#### Persistence Layer

The `pkg/task.go` module provides:

- `SaveToFile()`: Serializes task list to JSON
- `ReadFromFile()`: Deserializes JSON to task list
- Automatic handling of missing or empty files

#### Command Processing

Each command is implemented as a separate module in the `commands/` directory:

- Input validation and error handling
- File I/O operations
- User feedback messages

### Error Handling

The application includes comprehensive error handling for:

- Invalid command arguments
- File read/write operations
- JSON serialization/deserialization
- Task not found scenarios
- Invalid ID formats

## License

This project is available for personal and educational use.

