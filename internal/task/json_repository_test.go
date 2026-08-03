package task

import (
	"path/filepath"
	"testing"
)

func TestJSONRepositorySaveAndLoad(t *testing.T) {
	filePath := filepath.Join(
		t.TempDir(),
		"tasks.json",
	)

	repository := NewJSONRepository(filePath)

	expectedTasks := []Task{
		{
			ID:        1,
			Title:     "Learn Go",
			Completed: false,
		},
		{
			ID:        2,
			Title:     "Write tests",
			Completed: true,
		},
	}

	err := repository.Save(expectedTasks)
	if err != nil {
		t.Fatalf("expected no error saving tasks, got %v", err)
	}

	gotTasks, err := repository.Load()
	if err != nil {
		t.Fatalf("expected no error loading tasks, got %v", err)
	}

	if len(gotTasks) != len(expectedTasks) {
		t.Fatalf(
			"expected %d tasks, got %d",
			len(expectedTasks),
			len(gotTasks),
		)
	}

	for i := range expectedTasks {
		if gotTasks[i] != expectedTasks[i] {
			t.Errorf(
				"task %d does not match: expected %+v, got %+v",
				i,
				expectedTasks[i],
				gotTasks[i],
			)
		}
	}
}
