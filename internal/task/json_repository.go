package task

import (
	"encoding/json"
	"errors"
	"os"
)

type JSONRepository struct {
	filepath string
}

func NewJSONRepository(filePath string) *JSONRepository {
	return &JSONRepository{
		filepath: filePath,
	}

}

var _ Repository = (*JSONRepository)(nil)

func (r *JSONRepository) readTasks() ([]Task, error) {
	data, err := os.ReadFile(r.filepath)

	if err != nil {
		if os.IsNotExist(err) {
			return []Task{}, nil
		}
		return nil, err
	}

	var tasks []Task

	if err := json.Unmarshal(data, &tasks); err != nil {
		return nil, err
	}
	return tasks, nil
}

func(r *JSONRepository) WriteTasks(tasks []Task) error{
	data, err := json.MarshalIndent(tasks,"", " ")

	if err != nil{
		return err
	}

	return  os.WriteFile(r.filepath, data, 0644)
}

func (r *JSONRepository) Create(task Task) (Task, error) {
	tasks, err := r.readTasks()

	if err != nil {
		return Task{}, err
	}

	nextID := 1

	for _, currentTask := range tasks {
		if currentTask.ID >= nextID {
			nextID = currentTask.ID + 1
		}
	}

	task.ID = nextID
	task.Completed = false
	tasks = append(tasks, task)

	if err := r.WriteTasks(tasks); err != nil {
		return Task{}, err
	}

	return task, nil

}

func (r *JSONRepository) List() ([]Task, error) {
	return r.readTasks()

}

func (r *JSONRepository) Get(id int) (Task, error) {
	tasks, err := r.readTasks()
	if err != nil {
		return Task{}, err
	}

	for _, task := range tasks {
		if task.ID == id {
			return task, nil
		}
	}

	return Task{}, errors.New("Task not found")
}

func (r *JSONRepository) Update(task Task) error {
	tasks, err := r.readTasks()

	if err != nil {
		return err
	}

	for i := range tasks {
		if tasks[i].ID == task.ID {
			tasks[i] = task
			return r.WriteTasks(tasks)
		}
	}
	return errors.New("task not found")
}

func (r *JSONRepository) Delete(id int) error {
	tasks, err := r.readTasks()

	if err != nil {
		return  err
	}
	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			return r.WriteTasks(tasks)
		}
	}

	return errors.New("task not found")
}
