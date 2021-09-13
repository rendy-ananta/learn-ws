package repository

import (
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"testing"
	"time"
)

func TestTaskDb_GetAll(t *testing.T) {
	db := TaskDb{}

	all, err := db.GetAll()

	if err != nil {
		t.Errorf(err.Error())
	}

	fmt.Printf("task count %d \n", len(all))
}

func TestTaskDb_Find(t *testing.T) {
	db := TaskDb{}

	task := Task{
		Id:          uuid.NewString(),
		Name:        "Test Task",
		Description: sql.NullString{String: "Test task description", Valid: true},
		Status:      StatusOnProgress,
		DueDate:     sql.NullTime{Time: time.Now().Add(time.Hour * 12), Valid: true},
	}

	if err := db.Create(task); err != nil {
		t.Errorf("cannot create task: %v", err.Error())
	}

	result, err := db.Find(task.Id)

	if err != nil {
		t.Errorf("cannot get task: %v", err.Error())
	}

	if result.Name != task.Name {
		t.Errorf("name not match")
	}

	if result.Description != task.Description {
		t.Error(task.Description.String, ":", result.Description.String)
		t.Errorf("description not match")
	}

	if result.Status != task.Status {
		t.Errorf("status not match")
	}

	// Disable checking time due to postgresql limitation
	// of storing timezone.
	// ref:
	// - https://stackoverflow.com/a/28218103
	// - https://dba.stackexchange.com/a/164208

	//if result.DueDate.Time != task.DueDate.Time {
	//	fmt.Println(result.DueDate.Time.Unix())
	//	fmt.Println(task.DueDate.Time.Unix())
	//	t.Errorf("due_date not match")
	//}
}

func TestTaskDb_Update(t *testing.T) {
	db := TaskDb{}

	task := Task{
		Id:          uuid.NewString(),
		Name:        "Test Task Update",
		Description: sql.NullString{String: "Test task description", Valid: true},
		Status:      StatusOnProgress,
		DueDate:     sql.NullTime{Time: time.Now().Add(time.Hour * 12), Valid: true},
	}

	if err := db.Create(task); err != nil {
		t.Errorf("cannot create task: %v", err.Error())
	}

	task.Status = StatusDone

	if err := db.Update(task); err != nil {
		t.Errorf("cannot update task: %v", err)
	}

	result, err := db.Find(task.Id)

	if err != nil {
		t.Errorf("cannot get task: %v", err.Error())
	}

	if result.Status != StatusDone {
		t.Errorf("status does not match")
	}
}

func TestTaskDb_Delete(t *testing.T) {
	db := TaskDb{}

	task := Task{
		Id:          uuid.NewString(),
		Name:        "Test Task Update",
		Description: sql.NullString{String: "Test task description", Valid: true},
		Status:      StatusOnProgress,
		DueDate:     sql.NullTime{Time: time.Now().Add(time.Hour * 12), Valid: true},
	}

	if err := db.Create(task); err != nil {
		t.Errorf("cannot create task: %v", err.Error())
	}

	task.Status = StatusDone

	if err := db.Delete(task); err != nil {
		t.Errorf("cannot delete task: %v", err)
	}

	result, err := db.Find(task.Id)

	if err == nil {
		t.Errorf("finding task that have been deleted, should throw an error")
	}

	if result != (Task{}) {
		t.Errorf("task should be an empty task, since it has been deleted")
	}
}
