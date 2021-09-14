package repository

import (
	"fmt"
	"github.com/google/uuid"
	"web-svc/db/sql"

	_ "github.com/lib/pq"
)

type TaskDb struct{}

var client = sql.NewClient()

func (t TaskDb) GetAll() ([]Task, error) {
	var tasks []Task

	if err := client.Select(&tasks, "select * from tasks"); err != nil {
		return nil, fmt.Errorf("cannot execute statement: %v", err)
	}

	return tasks, nil
}

func (t TaskDb) Find(id string) (Task, error) {
	task := Task{}

	if err := client.Get(&task, client.Rebind("select * from tasks where id = ?"), id); err != nil {
		return Task{}, fmt.Errorf("cannot execute statement: %v", err)
	}

	return task, nil
}

func (t TaskDb) Create(task Task) error {
	if task.Id == "" {
		task.Id = uuid.NewString()
	}

	query, err := client.Prepare(client.Rebind("insert into tasks (id, name, description, due_date, status) values (?, ?, ?, ?, ?)"))

	if err != nil {
		return fmt.Errorf("cannot prepare statement: %v", err)
	}

	exec, err := query.Exec(task.Id, task.Name, task.Description, task.DueDate, task.Status)

	if affectedRow, err := exec.RowsAffected(); affectedRow == 0 && err != nil {
		return fmt.Errorf("cannot insert item: %v", err)
	}

	return nil
}

func (t TaskDb) Update(task Task) error {
	query, err := client.Prepare("update tasks set name = $1, description = $2, due_date = $3, status = $4 where id = $5")

	if err != nil {
		return fmt.Errorf("cannot prepare statement: %v", err)
	}

	exec, err := query.Exec(task.Name, task.Description, task.DueDate, task.Status, task.Id)

	if affectedRow, err := exec.RowsAffected(); affectedRow == 0 && err != nil {
		return fmt.Errorf("cannot insert item: %v", err)
	}

	return nil
}

func (t TaskDb) Delete(id string) error {
	query, err := client.Prepare("delete from tasks where id = $1")

	if err != nil {
		return fmt.Errorf("cannot prepare statement: %v", err)
	}

	exec, err := query.Exec(id)

	if affectedRow, err := exec.RowsAffected(); affectedRow == 0 && err != nil {
		return fmt.Errorf("cannot delete item: %v", err)
	}

	return nil
}

func (t TaskDb) UpdateStatus(id string, newStatus string) error {
	query, err := client.Prepare("update tasks set status = $1 where id = $2")

	if err != nil {
		return fmt.Errorf("cannot prepare statement %v: ", err)
	}

	exec, err := query.Exec(newStatus, id)

	if affectedRow, err := exec.RowsAffected(); affectedRow == 0 && err != nil {
		return fmt.Errorf("cannot update task status")
	}

	return nil
}
