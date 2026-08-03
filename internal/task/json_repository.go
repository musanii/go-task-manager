package task

import (
	"encoding/json"
	"os"
)

type JSONRepository struct {
	filepath string
}

func NewJSONRepository(filePath string) *JSONRepository {
	return &JSONRepository{
		filepath: filePath,
	}

}

var _ Repository = (*JSONRepository)(nil)

func (r *JSONRepository) Save(tasks []Task) error {
	data, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		return err
	}
	return os.WriteFile(r.filepath, data, 0644)
}

func (r *JSONRepository) Load() ([]Task, error) {
	data, err := os.ReadFile(r.filepath)

	if err != nil {
		if os.IsNotExist(err) {
			return []Task{}, nil
		}
		return nil, err
	}
	var tasks []Task

	err = json.Unmarshal(data, &tasks)
	if err != nil {
		return nil, err
	}
	return tasks, nil

}
