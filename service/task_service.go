package service

import (
	"context"

	"github.com/Revprm/go-fp-pbkk/dto"
	"github.com/Revprm/go-fp-pbkk/entity"
	"github.com/Revprm/go-fp-pbkk/repository"
	"github.com/google/uuid"
)

type (
	TaskService interface {
		CreateTask(ctx context.Context, req dto.TaskCreateRequest, userID string) (dto.TaskResponse, error)
		GetTaskByID(ctx context.Context, taskID uuid.UUID) (dto.TaskResponse, error)
		UpdateTask(ctx context.Context, taskID uuid.UUID, userID string, req dto.TaskUpdateRequest) (dto.TaskUpdateResponse, error)
		DeleteTask(ctx context.Context, taskID uuid.UUID, userID string) error
		GetTasksWithPagination(ctx context.Context, userID string, req dto.PaginationRequest) (dto.TaskPaginationResponse, error)
	}

	taskService struct {
		taskRepo repository.TaskRepository
	}
)

func NewTaskService(taskRepo repository.TaskRepository) TaskService {
	return &taskService{
		taskRepo: taskRepo,
	}
}

func (s *taskService) CreateTask(ctx context.Context, req dto.TaskCreateRequest, userId string) (dto.TaskResponse, error) {
	task := entity.Task{
		Title:       req.Title,
		Description: req.Description,
		Status:      req.Status,
		DueDate:     req.DueDate,
		UserID:      userId,
	}

	createdTask, err := s.taskRepo.CreateTask(ctx, nil, &task)
	if err != nil {
		return dto.TaskResponse{}, dto.ErrCreateTask
	}

	return dto.TaskResponse{
		ID:          createdTask.ID,
		Title:       createdTask.Title,
		Description: createdTask.Description,
		Status:      createdTask.Status,
		DueDate:     createdTask.DueDate,
		UserID:      createdTask.UserID,
	}, nil
}

func (s *taskService) GetTaskByID(ctx context.Context, taskID uuid.UUID) (dto.TaskResponse, error) {
	task, err := s.taskRepo.GetTaskByID(ctx, nil, taskID)
	if err != nil {
		return dto.TaskResponse{}, dto.ErrGetTask
	}

	return dto.TaskResponse{
		ID:          task.ID,
		Title:       task.Title,
		Description: task.Description,
		Status:      task.Status,
		DueDate:     task.DueDate,
		UserID:      task.UserID,
		User:        task.User.Name,
		CreatedAt:   task.CreatedAt,
		UpdatedAt:   task.UpdatedAt,
	}, nil
}

func (s *taskService) UpdateTask(ctx context.Context, taskID uuid.UUID, userID string, req dto.TaskUpdateRequest) (dto.TaskUpdateResponse, error) {
	existingTask, err := s.taskRepo.GetTaskByID(ctx, nil, taskID)
	if err != nil {
		return dto.TaskUpdateResponse{}, dto.ErrTaskNotFound
	}

	if existingTask.UserID != userID {
		return dto.TaskUpdateResponse{}, dto.ErrUserNotAllowed
	}

	task := &entity.Task{
		Title:       req.Title,
		Description: req.Description,
		Status:      req.Status,
		DueDate:     req.DueDate,
	}

	updatedTask, err := s.taskRepo.UpdateTask(ctx, nil, taskID, task)
	if err != nil {
		return dto.TaskUpdateResponse{}, dto.ErrUpdateTask
	}

	return dto.TaskUpdateResponse{
		ID:          updatedTask.ID,
		Title:       updatedTask.Title,
		Description: updatedTask.Description,
		Status:      updatedTask.Status,
		DueDate:     updatedTask.DueDate,
		UserID:      updatedTask.UserID,
		CreatedAt:   updatedTask.CreatedAt,
		UpdatedAt:   updatedTask.UpdatedAt,
	}, nil
}

func (s *taskService) DeleteTask(ctx context.Context, taskID uuid.UUID, userID string) error {
	existingTask, err := s.taskRepo.GetTaskByID(ctx, nil, taskID)
	if err != nil {
		return dto.ErrTaskNotFound
	}

	if existingTask.UserID != userID {
		return dto.ErrUserNotAllowed
	}

	if err := s.taskRepo.DeleteTask(ctx, nil, taskID); err != nil {
		return dto.ErrDeleteTask
	}

	return nil
}

func (s *taskService) GetTasksWithPagination(ctx context.Context, userID string, req dto.PaginationRequest) (dto.TaskPaginationResponse, error) {
	result, err := s.taskRepo.GetTasksWithPagination(ctx, nil, userID, req)
	if err != nil {
		return dto.TaskPaginationResponse{}, dto.ErrGetTasks
	}

	tasks := make([]dto.TaskResponse, len(result.Tasks))
	for i, task := range result.Tasks {
		tasks[i] = dto.TaskResponse{
			ID:          task.ID,
			Title:       task.Title,
			Description: task.Description,
			Status:      task.Status,
			DueDate:     task.DueDate,
			UserID:      task.UserID,
			User:        task.User.Name,
			CreatedAt:   task.CreatedAt,
			UpdatedAt:   task.UpdatedAt,
		}
	}

	return dto.TaskPaginationResponse{
		Data: tasks,
		PaginationResponse: dto.PaginationResponse{
			Page:    result.Page,
			PerPage: result.PerPage,
			Count:   result.Count,
			MaxPage: result.MaxPage,
		},
	}, nil
}
