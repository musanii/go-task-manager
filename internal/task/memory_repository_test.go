package task

import "testing"

type memoryRepository struct {
	tasks []Task
}

func (r *memoryRepository) Load() ([]Task, error) {
	return r.tasks, nil
}

func (r *memoryRepository) Save(tasks []Task) error {
	r.tasks = tasks
	return nil
}

func newTestService(t *testing.T) *Service {
	t.Helper()

	service, err := NewService(&memoryRepository{})
	if err != nil {
		t.Fatalf("failed to create service: %v", err)
	}

	return service
}
