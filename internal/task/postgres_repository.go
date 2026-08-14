package task

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
)

type PostgresRepository struct {
	db *pgx.Conn
}

func NewPostgresRepository(conn *pgx.Conn) *PostgresRepository {
	return &PostgresRepository{
		db: conn,
	}
}



func (r *PostgresRepository) Create(task Task) (Task, error) {
	ctx := context.Background()

	row := r.db.QueryRow(
		ctx,
		`INSERT INTO tasks (title,completed)
		VALUES ($1,$2)
		RETURNING id`,
		task.Title,
		task.Completed,
	)

	if err := row.Scan(&task.ID); err != nil {
		return Task{}, err

	}
	return task, nil

}

func (r *PostgresRepository) List() ([]Task, error) {
	ctx := context.Background()

	rows, err := r.db.Query(
		ctx,
		`SELECT id, title, completed FROM tasks ORDER BY id`,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var tasks []Task

	for rows.Next(){
		var task Task

		if err := rows.Scan(
			&task.ID,
			&task.Title,
			&task.Completed,

		); err != nil {
			return nil, err
		}
		tasks = append(tasks,task)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return tasks, nil
}

func (r *PostgresRepository) Get(id int) (Task, error) {
	ctx := context.Background()

	var task Task

	err := r.db.QueryRow(
		ctx,
		`SELECT id,title,completed FROM tasks where id=$1`,
		id,
	).Scan(
		&task.ID,
		&task.Title,
		&task.Completed,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return Task{}, errors.New("task not found")
		}
		return Task{}, err
	}
	return task, nil
}

func(r *PostgresRepository)Update(task Task) error {
	ctx := context.Background()

	result, err := r.db.Exec(
		ctx,
		`UPDATE tasks SET title = $1, completed = $2 WHERE id = $3`,
		task.Title,
		task.Completed,
	)

	if err != nil {
		return err
	}

	if result.RowsAffected()==0{
		return errors.New("task not found")
	}

	return nil
}

func (r *PostgresRepository) Delete(id int) error{
	ctx := context.Background()
	result, err := r.db.Exec(
		ctx,
		"DELETE FROM tasks WHERE id = $1",
		id,
	)

	if err != nil{
		return err
	}

	if result.RowsAffected() == 0 {
		return errors.New("task not found")
	}

	return nil
}
