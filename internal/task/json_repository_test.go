package task

import (
	"path/filepath"
	"testing"
)



func TestJSONRepositoryList(t *testing.T) {
	filePath := filepath.Join(
		t.TempDir(),
		"tasks.json",
	)

	repository := NewJSONRepository(filePath)
	_, err := repository.Create(Task{
		Title: "Learn Go",
	})
	if err != nil {
		t.Fatalf("expected no error creating tasks, got %v", err)
	}

	_, err = repository.Create(Task{
		Title: "Write tests",
		Completed: true,
	})
	if err != nil {
		t.Fatalf("expected no error creating tasks, got %v", err)
	}


	got, err := repository.List()
	if err != nil {
		t.Fatalf("expected no error listing tasks, got %v", err)
	}

	if len(got) != 2{
		t.Fatalf("expected 2  tasks, got %d", len(got))
	}

	if got[0].Title != "Learn Go" {
		t.Errorf("expected first task %q, got %q",
			"Learn Go", got[0].Title)
	}

	if got[1].Title != "Write tests" {
		t.Errorf("expected second task %q, got %q",
			"Write tests", got[1].Title)
	}
}

func TestJSONRepositoryCreate(t *testing.T){
	filepath := filepath.Join(
		t.TempDir(),
		"tasks.json",
	)

	repository := NewJSONRepository(filepath)

	task :=Task{
		Title: "Learn Go",
	}

	got, err := repository.Create(task)
	if err != nil{
		t.Fatalf("expected no error creating task, got %v", err)
	}

	if got.ID != 1{
		t.Fatalf("expected ID 1, got %d", got.ID)
	}

	if got.Title != "Learn Go"{
		t.Errorf("expected title Learn Go, got %qq", got.Title)
	}

	if got.Completed {
		t.Errorf("expected task to be incomplete")
	}

}
