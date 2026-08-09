package task

import (
	"context"

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

func (r *PostgresRepository) Save(tasks []Task) error {
	ctx := context.Background()

	_, err := r.db.Exec(ctx, "DELETE FROM tasks")

	if err != nil {
		return err
	}

	for _, task := range tasks {
		_, err := r.db.Exec(
			ctx,
			"INSERT INTO tasks (id,title,completed) VALUES ($1,$2,$3)",
			task.ID,
			task.Title,
			task.Completed,
		)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *PostgresRepository) Load() ([]Task, error) {
	ctx := context.Background()

	rows, err := r.db.Query(
		ctx,
		"SELECT id, title, completed FROM tasks ORDER BY id",
	)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var tasks []Task

	for rows.Next() {
		var task Task

		err := rows.Scan(
			&task.ID,
			&task.Title,
			&task.Completed,
		)

		if err != nil {
			return nil, err
		}

		tasks = append(tasks, task)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil

}
