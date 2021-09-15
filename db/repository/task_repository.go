package repository

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"log"
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

func (task Task) MarshalBinary() (data []byte, err error) {
	bytes, err := json.Marshal(task)
	return bytes, err
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
	Update(task Task) error
	Delete(task Task) error
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

func (TaskRepository) Find(id string) (Task, error) {
	var task Task

	cachedContent, err := TaskCache{}.Find(id)

	if cachedContent == (Task{}) && err != nil {
		log.Printf("cache not found, querying data...")

		queryResult, err := TaskDb{}.Find(id)

		if err != nil {
			return Task{}, fmt.Errorf("cannot querying data: %v", err)
		}

		task = queryResult

		// set the cache
		go func() {
			err := TaskCache{}.Set(task)
			if err != nil {
				log.Fatalf("error set cache: %v", err)
			}
		}()
	} else {
		log.Printf("using data from cache")
		task = cachedContent
	}

	return task, nil
}

func (TaskRepository) Create(task Task) (Task, error) {
	if task.Id == "" {
		task.Id = uuid.NewString()
	}

	if task.Status == "" { // default value
		task.Status = StatusNotStarted
	}

	if err := (TaskDb{}).Create(task); err != nil {
		return Task{}, fmt.Errorf("cannot create task: %v", err)
	}

	item, err := TaskDb{}.Find(task.Id)
	if err != nil {
		return Task{}, fmt.Errorf("cannot fetch created task: %v", err)
	}

	// set the cache
	go func() {
		err := TaskCache{}.Set(task)
		if err != nil {
			log.Fatalf("error set cache: %v", err)
		}
	}()

	return item, nil
}

func (TaskRepository) Update(task Task) (Task, error) {
	if err := (TaskDb{}).Update(task); err != nil {
		return Task{}, fmt.Errorf("cannot update task: %v", err)
	}

	item, err := TaskDb{}.Find(task.Id)
	if err != nil {
		return Task{}, fmt.Errorf("cannot fetch updated task: %v", err)
	}

	// set the cache
	go func() {
		err := TaskCache{}.Set(task)
		if err != nil {
			log.Fatalf("error set cache: %v", err)
		}
	}()

	return item, nil
}

func (t TaskRepository) Delete(id string) error {
	if err := (TaskDb{}).Delete(id); err != nil {
		return fmt.Errorf("cannot delete task: %v", err)
	}

	return nil
}

func (TaskRepository) UpdateStatus(id string, newStatus string) (Task, error) {
	if err := (TaskDb{}).UpdateStatus(id, newStatus); err != nil {
		return Task{}, fmt.Errorf("cannot update task status: %v", err)
	}

	task, err := TaskDb{}.Find(id)
	if err != nil {
		return Task{}, fmt.Errorf("cannot fetch updated task status: %v", err)
	}

	// set the cache
	go func() {
		err := TaskCache{}.Set(task)
		if err != nil {
			log.Fatalf("error set cache: %v", err)
		}
	}()

	return task, nil
}
