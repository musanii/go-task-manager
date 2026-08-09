package task

import "fmt"

func NewRepository(storageType string, taskFile string) (Repository, error) {
	switch storageType {
	case "json":
		return NewJSONRepository(taskFile), nil

	case "postgres":
		return nil, fmt.Errorf("PostgreSQL storage is not implimented yet.")

	default:
		return nil, fmt.Errorf("Unsupported storage type: %s\n", storageType)

	}

}
