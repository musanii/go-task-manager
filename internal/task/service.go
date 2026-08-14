package task

import "errors"

type Service struct {
	repository Repository
	
}

func NewService(repository Repository) (*Service, error) {

	return &Service{
		repository: repository,
	}, nil

}

func (s *Service) Create(title string) (Task, error) {
	if title == "" {
		return Task{}, errors.New("task title is required")
	}
	task := Task{
		
		Title:     title,
		Completed: false,
	}
	return s.repository.Create(task)
}

func (s *Service) List() ([]Task, error) {
	return s.repository.List()
}

func (s *Service) Complete(id int) error {
	task, err := s.repository.Get(id)
	if err != nil {
		return err
	}

	task.Completed= true
	return s.repository.Update(task)

}

func (s *Service) Delete(id int) error {
	return s.repository.Delete(id)
}

func (s *Service) Update(id int, title string) error {
	if title == "" {
		return errors.New("task title is required")
	}
	task, err := s.repository.Get(id)
	if err != nil {
		return err
	}
	task.Title = title
	return s.repository.Update(task)

	
}
