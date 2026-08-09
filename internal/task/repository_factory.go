package task

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func NewRepository(
	storageType string,
	taskFile string,
	databaseURL string,
) (Repository, error) {
	switch storageType {
	case "json":
		return NewJSONRepository(taskFile), nil

	case "postgres":
		conn, err := pgx.Connect(context.Background(), databaseURL)
		if err != nil {
			return nil, err
		}
		return NewPostgresRepository(conn),nil

	default:
		return nil, fmt.Errorf("Unsupported storage type: %s\n", storageType)

	}

}
