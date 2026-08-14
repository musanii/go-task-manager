package task

import "testing"

func TestServiceCreate(t *testing.T) {

	tests := []struct {
		name    string
		title   string
		wantErr string
	}{
		{
			name:    "valid title",
			title:   "Learn Go",
			wantErr: "",
		},
		{
			name:    "empty title",
			title:   "",
			wantErr: "task title is required",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			service := newTestService(t)
			_, err := service.Create(tt.title)
			if tt.wantErr == "" {
				if err != nil {
					t.Fatalf("expected no error, got %v", err)
				}
			} else {
				if err == nil {
					t.Fatalf("expected error %q, got nil", tt.wantErr)
				}

				if err.Error() != tt.wantErr {
					t.Fatalf("expected error %q. got %q", tt.wantErr, err.Error())
				}
			}
		})
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
	tasks, err := service.List()

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

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

	tasks, err := service.List()

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

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

	tasks, err := service.List()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

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

func TestServiceUpdate(t *testing.T) {
	service := newTestService(t)

	createdTask, err := service.Create("Learn Go")

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	err = service.Update(createdTask.ID, "Learn Go properly")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	tasks, err := service.List()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(tasks) != 1 {
		t.Fatalf("expected 1 task, got %d", len(tasks))
	}

	if tasks[0].Title != "Learn Go properly" {
		t.Errorf(
			"expected title %q, got %q",
			"Learn Go properly",
			tasks[0].Title,
		)
	}
}

func TestServiceSearch(t *testing.T) {

	service := newTestService(t)
	_, err := service.Create("Learn Go")
	if err != nil {
		t.Fatal(err)
	}

	_, err = service.Create("Write PostgreSQL tests")
	if err != nil {
		t.Fatal(err)
	}

	_, err = service.Create("Build CLI")
	if err != nil {
		t.Fatal(err)
	}

	tasks, err := service.Search("Go")
	if err != nil {
		t.Fatal(err)
	}

	if len(tasks) != 1 {
		t.Fatalf("expected 1 tasks got %d", len(tasks))
	}

	if tasks[0].Title != "Learn Go" {
		t.Errorf(
			"expected %q, got %q",
			"Learn Go",
			tasks[0].Title,
		)
	}

}

func TestServiceReturnsEmptyWhenNoMatch(t *testing.T) {
	service := newTestService(t)

	_, err := service.Create("Learn Go")

	if err != nil {
		t.Fatal(err)
	}

	_, err = service.Create("Write PostgreSQL tests")

	if err != nil {
		t.Fatal(err)
	}

	tasks, err := service.Search("Ruby")
	if err != nil {
		t.Fatal(err)
	}

	if len(tasks) != 0 {
		t.Fatalf("expected 0 tasks, got %d", len(tasks))
	}
}

func TestServiceSearchIsCaseInsensitive(t *testing.T) {

	service := newTestService(t)

	_,err := service.Create("Learn Go")
	if err != nil {
		t.Fatal(err)
	}

	tasks,err := service.Search("go")
	 if err != nil{
		t.Fatal(err)
	 }

	 if len(tasks) != 1 {
		t.Fatalf("expected 1 task,  got %d", len(tasks))
	 }

	 if tasks[0].Title != "Learn Go"{
		t.Errorf(
			"expected %q, got %q",
			"Learn Go",
			tasks[0].Title,
		)
	 }
}

func TestServiceRejectsEmptyKeyword(t *testing.T)  {
	service := newTestService(t)

	_,err := service.Create(" Learn Go")

	if err != nil {
		t.Fatal(err)
	}

	_, err = service.Search("  ")

	if err == nil {
		t.Fatal("expected search keyword error, got nil")
	}

	expected := "search keyword is required"
	if err.Error() != expected {
		t.Fatalf(
			"expected error %q, got %q",
			expected,
			err.Error(),
		)
	}

	

	
}
