package repository

import (
	"database/sql"
	"fmt"
	"time"
)

type Task struct {
	Id          string         `db:"id"`
	Name        string         `db:"name"`
	Description sql.NullString `db:"description"`
	Status      string         `db:"status"`
	CreatedAt   time.Time      `db:"created_at"`
	DueDate     sql.NullTime   `db:"due_date"`
}

const (
	StatusOnProgress = "On Progress"
	StatusDone       = "Done"
	StatusNotStarted = "Not Started"
	StatusClosed     = "Closed"
)

type TaskRepositoryContract interface {
	GetAll() ([]Task, error)
	Find(id string) (Task, error)
	Create(task Task) error
	Update(id string, task Task) error
	Delete(id string, task Task) error
	UpdateStatus(id string, newStatus string) error
}

type TaskRepository struct{}

func (TaskRepository) GetAll() ([]Task, error) {
	query, err := TaskDb{}.GetAll()

	if err != nil {
		return nil, fmt.Errorf("cannot querying data %v", err)
	}

	return query, nil
}

func (TaskRepository) Find(id string) (*Task, error) {
	panic("implement me")
}

func (TaskRepository) Create(task Task) (*Task, error) {
	panic("implement me")
}

func (TaskRepository) Update(id string, task Task) (*Task, error) {
	panic("implement me")
}

func (TaskRepository) Delete(id string, task Task) (*Task, error) {
	panic("implement me")
}

func (TaskRepository) UpdateStatus(id string, newStatus string) (*Task, error) {
	panic("implement me")
}
