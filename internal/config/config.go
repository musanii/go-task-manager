package config

import "os"

type Config struct {
	TaskFile string
	StorageType string
}

func Load() Config {
	file := os.Getenv("TASK_FILE")

	

	if file == ""{
		file = "tasks.json"
	}

	storageType := os.Getenv("STORAGE")

	if storageType == ""{
		storageType="json"
	}

	

	return Config{
		TaskFile: file,
		StorageType: storageType,
	}
}