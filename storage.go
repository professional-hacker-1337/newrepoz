package main

type Storage interface {
	Add(task *Task) error
	Get(id int) (*Task, error)
	List(status string) ([]*Task, error)
	Update(task *Task) error
	Delete(id int) error
}
