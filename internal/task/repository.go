package task

type Repository interface {
	Create(Task) (Task, error)
	List() ([]Task, error)
	Get(id int) (Task, error)
	Update(Task) error
	Delete(id int) error
}
