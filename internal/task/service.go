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

func (s *Service) Complete(id int) error {
	for i := range s.tasks {
		if s.tasks[i].ID == id {
			s.tasks[i].Completed = true
			return nil
		}
	}
	return errors.New("task not found")

}

func (s *Service) Delete(id int) error {
	for i, task := range s.tasks {
		if task.ID == id {
			s.tasks = append(
				s.tasks[:i],
				s.tasks[i+1:]...,
			)

			return nil
		}
	}
	return errors.New("task not found")
}
