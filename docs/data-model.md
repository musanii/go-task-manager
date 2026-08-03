# Task Data Model

## Task

A task represents a piece of work managed by a single local user.

| Field     | Type    | Description                            |
| --------- | ------- | -------------------------------------- |
| ID        | Integer | Unique identifier for the task         |
| Title     | String  | Description of the task                |
| Completed | Boolean | Indicates whether the task is complete |

## Example

```json
{
  "id": 1,
  "title": "Learn Go",
  "completed": false
}
```

## Design Decisions

### Integer IDs

The MVP uses sequential integer IDs because they are simple for users to enter in CLI commands.

Example:

```bash
task complete 1
```

A UUID is unnecessary for a single-user local application.

### Boolean completion status

The MVP only needs two task states:

* Incomplete
* Complete

A boolean is sufficient.

A more complex state model, such as `pending`, `in_progress`, `blocked`, and `completed`, is outside the MVP.

### No timestamps yet

The MVP does not include creation or completion timestamps because no current feature requires them.

They may be added later if the application introduces task history, sorting, reporting, or due dates.
