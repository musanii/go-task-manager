package task

import "testing"

func TestServiceCreate(t *testing.T) {
	service := newTestService(t)
	got, err := service.Create("Learn Go")

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if got.ID != 1 {
		t.Errorf("expected ID 1 got %d", got.ID)
	}

	if got.Title != "Learn Go" {
		t.Errorf(
			"expected title %q, got %q",
			"Learn Go",
			got.Title,
		)
	}

	if got.Completed {
		t.Error("expected new task to be incomplete")
	}
}

func TestServiceCreateRejectsEmptyTitle(t *testing.T) {
	service := newTestService(t)

	_, err := service.Create("")
	if err == nil {
		t.Fatal("expected an error,got nil")
	}
	expected := "task title is required"
	if err.Error() != expected {
		t.Errorf(
			"expected error %q, got %q",
			expected,
			err.Error(),
		)
	}
}

func TestServiceList(t *testing.T) {
	service := newTestService(t)

	_, err := service.Create("Learn Go")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	_, err = service.Create("Write tests")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	tasks := service.List()

	if len(tasks) != 2 {
		t.Fatalf(
			"expected 2 tasks, got %d",
			len(tasks),
		)
	}

	if tasks[0].Title != "Learn Go" {
		t.Errorf(
			"expected first task %q, got %q",
			"Learn Go",
			tasks[0].Title,
		)
	}

	if tasks[1].Title != "Write tests" {
		t.Errorf(
			"expected second task %q, got %q",
			"Write tests",
			tasks[1].Title,
		)
	}
}

func TestServiceComplete(t *testing.T) {
	service := newTestService(t)

	createdTask, err := service.Create("Learn Go")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	err = service.Complete(createdTask.ID)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	tasks := service.List()

	if len(tasks) != 1 {
		t.Fatalf(
			"expected 1 task, got %d",
			len(tasks),
		)
	}
	if !tasks[0].Completed {
		t.Error("expected task to be completed")
	}
}

func TestServiceCompleteReturnsErrorForMissingTask(t *testing.T) {
	service := newTestService(t)

	err := service.Complete(99)

	if err == nil {
		t.Fatalf("expected an error, got nil")
	}

	expected := "task not found"

	if err.Error() != expected {
		t.Errorf(
			"expected error %q, got %q",
			expected,
			err.Error(),
		)
	}
}

func TestServiceDelete(t *testing.T) {
	service := newTestService(t)

	first, err := service.Create("Learn Go")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	_, err = service.Create("Write tests")
	if err != nil {
		t.Fatalf("expected no error got %v", err)
	}
	third, err := service.Create("Build CLI")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	err = service.Delete(2)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	tasks := service.List()

	if len(tasks) != 2 {
		t.Fatalf("expected 2 tasks, got %d", len(tasks))
	}
	if tasks[0].ID != first.ID {
		t.Errorf(
			"expected first remaining task ID %d, got %d",
			first.ID,
			tasks[0].ID,
		)
	}

	if tasks[1].ID != third.ID {
		t.Errorf(
			"expected second remaining task ID %d, got %d",
			third.ID,
			tasks[1].ID,
		)
	}
}

func TestServiceDeleteReturnsErrorForMissingTask(t *testing.T) {
	service := newTestService(t)

	err := service.Delete(99)
	if err == nil {
		t.Fatal("expected an error got nil")
	}

	expected := "task not found"

	if err.Error() != expected {
		t.Errorf(
			"expected error %q, got %q",
			expected,
			err.Error(),
		)
	}
}
