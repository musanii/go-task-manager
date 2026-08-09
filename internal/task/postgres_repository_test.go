package task

import (
	"context"
	"testing"

	"github.com/jackc/pgx/v5"
)

func TestPostgresRepositorySaveAndLoad(t *testing.T) {
	conn, err := pgx.Connect(
		context.Background(),
		"postgres://postgres:Osiron@1994@localhost:5432/task_manager_test",
	)

	if err != nil {
		t.Fatal(err)
	}

	defer conn.Close(context.Background())

	_, err = conn.Exec(context.Background(), "DELETE FROM TASKS")
	if err != nil {
		t.Fatal(err)
	}

	repo := NewPostgresRepository(conn)

	tasks := []Task{
		{
			ID:        1,
			Title:     "Learn PostgreSQL",
			Completed: false,
		},
	}

	err = repo.Save(tasks)
	if err != nil {
		t.Fatal(err)
	}
	got, err := repo.Load()
	if err != nil {
		t.Fatal(err)
	}

	if len(got) != 1 {
		t.Fatalf("expected 1 task got %d", len(got))
	}

	if got[0] != tasks[0] {
		t.Fatalf("expected %+v, got %+v", tasks[0], got[0])
	}

}
