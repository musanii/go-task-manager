package task

import "testing"

func TestNewRepository(t *testing.T) {
	tests := []struct {
		name        string
		storageType string
		wantErr     bool
	}{
		{
			name:        "json repository",
			storageType: "json",
			wantErr:     false,
		},
		{
			name:        "unsupported repository",
			storageType: "mysql",
			wantErr:     true,
		},
		{
			name:        "postgres repository",
			storageType: "postgres",
			wantErr:     true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo, err := NewRepository(
				tt.storageType,
				 "test-tasks.json",
				"postgres://postgres:test@localhost:5432/task_manager_test",
			)

			if (err != nil) != tt.wantErr {
				t.Fatalf(
					"expected error: %v, got: %v",
					tt.wantErr,
					err,
				)
			}
			if !tt.wantErr && repo == nil {
				t.Fatal("expected repository, got nil")
			}
		})
	}
}
