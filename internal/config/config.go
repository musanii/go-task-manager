package config

import "os"

type Config struct {
	TaskFile    string
	StorageType string
	DatabaseURL string
}

func Load() Config {
	file := os.Getenv("TASK_FILE")

	if file == "" {
		file = "tasks.json"
	}

	storageType := os.Getenv("STORAGE")

	if storageType == "" {
		storageType = "json"
	}
	databaseURL := os.Getenv("DATABASE_URL")

	if databaseURL == "" {
		databaseURL = "postgres://postgres:Osiron@1994@localhost:5432/task_manager"
	}

	return Config{
		TaskFile:    file,
		StorageType: storageType,
		DatabaseURL: databaseURL,
	}
}
