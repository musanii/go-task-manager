package task

import "errors"

type Service struct {
	nextID int
	tasks  []Task
}

func NewService() *Service {
	return &Service{
		nextID: 1,
	}
}

func (s *Service) Create(title string) (Task, error) {
	if title == "" {
		return Task{}, errors.New("task title is required")
	}
	task := Task{
		ID:        s.nextID,
		Title:     title,
		Completed: false,
	}
	s.tasks = append(s.tasks, task)
	s.nextID++
	return task, nil
}

func (s *Service) List() []Task {
	return s.tasks
}
