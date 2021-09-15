package server

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
	"web-svc/db/repository"
)

var taskRepository = repository.TaskRepository{}

type TaskManagerServerImpl struct {
	UnimplementedTaskManagerServer
}

func (t TaskManagerServerImpl) Create(ctx context.Context, request *CreateTaskRequest) (*Task, error) {
	response, err := taskRepository.Create(repository.Task{
		Name: request.Name,
		Description: sql.NullString{
			Valid:  request.Description != "",
			String: request.Description,
		},
		Status:    request.Status,
		CreatedAt: time.Time{}, // will be created automatically in database,
		DueDate: sql.NullTime{
			Valid: request.DueDate.IsValid(),
			Time:  request.DueDate.AsTime(),
		},
	})

	if err != nil {
		return nil, fmt.Errorf("error invoking `Create()` in task repository: %v", err)
	}

	return mapFromTaskRepository(response), nil
}

func (t TaskManagerServerImpl) Find(ctx context.Context, request *TaskIdRequest) (*Task, error) {
	response, err := taskRepository.Find(request.Id)

	if err != nil {
		return nil, fmt.Errorf("error invoking `Find()` in task repository: %v", err)
	}

	return mapFromTaskRepository(response), nil
}

func (t TaskManagerServerImpl) GetAll(ctx context.Context, empty *emptypb.Empty) (*TaskList, error) {
	list := new(TaskList)

	response, err := taskRepository.GetAll()

	if err != nil {
		return nil, fmt.Errorf("error invoking `GetAll()` in task repository: %v", err)
	}

	for i := 0; i < len(response); i++ {
		list.List = append(list.List, mapFromTaskRepository(response[i]))
	}

	return list, nil
}

func (t TaskManagerServerImpl) Update(ctx context.Context, request *UpdateTaskRequest) (*Task, error) {
	response, err := taskRepository.Update(repository.Task{
		Id:   request.Id,
		Name: request.Name,
		Description: sql.NullString{
			Valid:  request.Description != "",
			String: request.Description,
		},
		Status: request.Status,
		DueDate: sql.NullTime{
			Valid: request.DueDate.IsValid(),
			Time:  request.DueDate.AsTime(),
		},
	})

	if err != nil {
		return nil, fmt.Errorf("error invoking `Update()` in task repository: %v", err)
	}

	return mapFromTaskRepository(response), nil
}

func (t TaskManagerServerImpl) UpdateStatus(ctx context.Context, request *TaskStatusRequest) (*Task, error) {
	response, err := taskRepository.UpdateStatus(request.Id, request.Status)

	if err != nil {
		return nil, fmt.Errorf("error invoking `UpdateStatus() in task repository: %v`", err)
	}

	return mapFromTaskRepository(response), nil
}

func (t TaskManagerServerImpl) Delete(ctx context.Context, request *TaskIdRequest) (*emptypb.Empty, error) {
	if err := taskRepository.Delete(request.Id); err != nil {
		return nil, fmt.Errorf("error invoking `Delete() in task repository: %v`", err)
	}

	return &empty.Empty{}, nil
}

func mapFromTaskRepository(task repository.Task) *Task {
	var dueDate *timestamppb.Timestamp

	if task.DueDate.Valid {
		dueDate = timestamppb.New(task.DueDate.Time)
	}

	return &Task{
		Id:          task.Id,
		Name:        task.Name,
		Description: task.Description.String,
		Status:      task.Status,
		CreatedAt:   timestamppb.New(task.CreatedAt),
		DueDate:     dueDate,
	}
}
