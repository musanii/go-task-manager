# Go Task Manager

A command-line task management application built with Go.

## Problem

People need a simple way to create, view, complete, and manage tasks from the command line.

This project is being built to practise:

* Idiomatic Go
* System design
* Software architecture
* Testing
* Git and GitHub workflows
* Building maintainable software

## MVP Features

The first version will allow a user to:

* Add a task
* List tasks
* Mark a task as complete
* Delete a task

## Example Usage

```bash
task add "Learn Go"
task list
task complete 1
task delete 1
```

## Initial System Design

```text
User
  |
  | CLI command
  v
CLI Application
  |
  | task operation
  v
Task Service
  |
  | save or retrieve tasks
  v
Local JSON File
```

## Technology

* Go
* Go standard library
* JSON file storage
* Go's built-in testing package

## Project Status

🚧 In development
