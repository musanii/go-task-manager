package task

import "testing"

func TestServiceCreate(t *testing.T) {
	service := NewService()
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
	service := NewService()

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
	service := NewService()

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
