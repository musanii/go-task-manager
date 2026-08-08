package task

import "errors"

type Service struct {
	repository Repository
	nextID     int
	tasks      []Task
}

func NewService(repository Repository) (*Service, error) {

	tasks, err := repository.Load()
	if err != nil {
		return nil, err
	}

	nextID := 1

	for _, currentTask := range tasks {
		if currentTask.ID >= nextID {
			nextID = currentTask.ID + 1
		}
	}
	return &Service{
		repository: repository,
		nextID:     nextID,
		tasks:      tasks,
	}, nil

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

	if err := s.repository.Save(s.tasks); err != nil {
		return Task{}, err
	}
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

			if err := s.repository.Save(s.tasks); err != nil {
				return err
			}
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
			if err := s.repository.Save(s.tasks); err != nil {
				return err
			}

			return nil
		}
	}
	return errors.New("task not found")
}

func (s *Service) Update(id int, title string) error {
	if title == "" {
		return errors.New("task title is required")
	}
	tasks, err := s.repository.Load()
	if err != nil {
		return err
	}

	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Title = title
			return s.repository.Save(tasks)
		}
	}
	return errors.New("task not found")
}
