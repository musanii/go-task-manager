package task

import (
	"errors"
	"testing"
)

type memoryRepository struct {
	tasks []Task
}



func newTestService(t *testing.T) *Service {
	t.Helper()

	service, err := NewService(&memoryRepository{})
	if err != nil {
		t.Fatalf("failed to create service: %v", err)
	}

	return service
}

func (r *memoryRepository) Create(task Task) (Task, error) {
	nextID := 1

	for _, currentTask := range r.tasks {
		if currentTask.ID >= nextID {
			nextID = currentTask.ID + 1
		}
	}

	task.ID = nextID
	task.Completed = false
	r.tasks = append(r.tasks, task)
	return task, nil

}


func(r *memoryRepository)List()([]Task,error){
	return r.tasks,nil
}

func (r *memoryRepository) Get(id int) (Task, error) {
	for _,task := range r.tasks{
		if task.ID ==id{
			return task, nil
		}
	}

	return Task{}, errors.New("task not found")
}


func(r *memoryRepository) Update(task Task) error{
	for i := range r.tasks{
		if r.tasks[i].ID == task.ID{
			r.tasks[i] = task
			return nil
		}
	}
	return errors.New("task not found")
}

func(r * memoryRepository) Delete(id int) error{
	for i, task := range r.tasks{
		if task.ID == id {
			r.tasks = append(r.tasks[:i],r.tasks[i+1:]... )
			return nil
		}
	}

	return errors.New("task not found")
}

