package task

import (
	"context"
	"testing"

	"github.com/jackc/pgx/v5"
)

func TestPostgresRepositoryList(t *testing.T) {
	conn, err := pgx.Connect(
		context.Background(),
		"postgres://postgres:Osiron@1994@localhost:5432/task_manager_test",
	)
	if err != nil {
		t.Fatal(err)
	}

	defer conn.Close(context.Background())

	_, err = conn.Exec(context.Background(), "DELETE FROM tasks")
	if err != nil {
		t.Fatal(err)
	}

	repo := NewPostgresRepository(conn)

	_, err = repo.Create(Task{
		Title: "Learn PostgreSQL",
	})
	if err != nil {
		t.Fatal(err)
	}

	_, err = repo.Create(Task{
		Title:     "Write tests",
		Completed: true,
	})
	if err != nil {
		t.Fatal(err)
	}

	got, err := repo.List()
	if err != nil {
		t.Fatal(err)
	}

	if len(got) != 2 {
		t.Fatalf("expected 2 tasks, got %d", len(got))
	}

	if got[0].Title != "Learn PostgreSQL" {
		t.Errorf(
			"expected first task %q, got %q",
			"Learn PostgreSQL",
			got[0].Title,
		)
	}

	if got[1].Title != "Write tests" {
		t.Errorf(
			"expected second task %q, got %q",
			"Write tests",
			got[1].Title,
		)
	}

	if !got[1].Completed {
		t.Error("expected second task to be completed")
	}
}

func TestPostgresRepositoryCreate(t *testing.T){
	conn,err := pgx.Connect(
		context.Background(),
		"postgres://postgres:Osiron@1994@localhost:5432/task_manager_test",
	)

	if err != nil {
		t.Fatal(err)
	}

	defer conn.Close(context.Background())

	_,err = conn.Exec(context.Background(),"DELETE FROM tasks")
	if err != nil{
		t.Fatal(err)
	}

	repo := NewPostgresRepository(conn)

	input := Task{
		Title: "Learn PostgreSQL",
		Completed: false,
	}

	got, err := repo.Create(input)
	if err != nil {
		t.Fatal(err)
	}

	if got.ID <= 0{
		t.Fatalf("expected ID 1, got %d", got.ID)
		
	}

	if got.Title != "Learn PostgreSQL"{
		t.Fatalf(
			"expected title %q, got %q",
			"Learn PostgreSQL",
			got.Title,
		)
	}

	if got.Completed{
		t.Fatal("expected task to be incomplete")
	}
}
