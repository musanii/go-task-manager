# Go Task Manager — Requirements

## 1. Problem Statement

Users need a lightweight way to manage personal tasks from the command line without requiring an account, internet connection, or external database.

The application should allow users to create and manage tasks using simple commands.

## 2. Target User

The initial target user is a single person using the application locally from a terminal.

The application does not currently support:

* Multiple users
* User accounts
* Authentication
* Shared task lists
* Remote access

## 3. MVP Functional Requirements

The first version must allow a user to:

### Add a task

The user can create a task by providing a title.

Example:

```bash
task add "Learn Go"
```

Expected result:

```text
Task created successfully.
```

### List tasks

The user can view all saved tasks.

Example:

```bash
task list
```

Expected result:

```text
ID  Status  Title
1   [ ]     Learn Go
2   [x]     Write tests
```

### Complete a task

The user can mark a task as completed using its ID.

Example:

```bash
task complete 1
```

Expected result:

```text
Task 1 marked as complete.
```

### Delete a task

The user can delete a task using its ID.

Example:

```bash
task delete 1
```

Expected result:

```text
Task 1 deleted.
```

## 4. Validation Rules

The application must:

* Reject empty task titles
* Reject invalid task IDs
* Return an error when a requested task does not exist
* Display useful error messages
* Avoid corrupting saved task data

## 5. Non-Functional Requirements

The application should:

* Run locally
* Work without an internet connection
* Start quickly
* Store data persistently
* Be easy to test
* Use the Go standard library where practical
* Have clear documentation
* Maintain a clean Git history

## 6. Out of Scope

The MVP will not include:

* A web interface
* A REST API
* Authentication
* Multiple users
* Task sharing
* Task reminders
* Notifications
* Categories
* Due dates
* PostgreSQL
* Docker
* Cloud deployment

These features may be added in later versions if they solve a real requirement.

## 7. Success Criteria

The MVP is complete when a user can:

1. Add a task
2. Close the application
3. Run the application again
4. View the previously saved task
5. Mark the task as complete
6. Delete the task

All core behaviour must be covered by automated tests.
