package dto

import (
	"errors"
	"time"

	"github.com/Revprm/go-fp-pbkk/entity"
	"github.com/google/uuid"
)

const (
	// Failed
	MESSAGE_FAILED_GET_TASK    = "failed to get task"
	MESSAGE_FAILED_CREATE_TASK = "failed to create task"
	MESSAGE_FAILED_UPDATE_TASK = "failed to update task"
	MESSAGE_FAILED_DELETE_TASK = "failed to delete task"
	MESSAGE_FAILED_GET_TASKS   = "failed to get tasks"

	// Success
	MESSAGE_SUCCESS_GET_TASK    = "success getting task"
	MESSAGE_SUCCESS_GET_TASKS   = "success getting tasks"
	MESSAGE_SUCCESS_CREATE_TASK = "success creating task"
	MESSAGE_SUCCESS_UPDATE_TASK = "success updating task"
	MESSAGE_SUCCESS_DELETE_TASK = "success deleting task"
)

var (
	ErrCreateTask     = errors.New("failed to create task")
	ErrGetTask        = errors.New("failed to get task")
	ErrUpdateTask     = errors.New("failed to update task")
	ErrDeleteTask     = errors.New("failed to delete task")
	ErrGetTasks       = errors.New("failed to get tasks")
	ErrUserNotAllowed = errors.New("user not allowed to perform this action")
	ErrTaskNotFound   = errors.New("task not found")
)

type (
	TaskCreateRequest struct {
		Title       string    `json:"title" form:"title"`
		Description string    `json:"description" form:"description"`
		Status      string    `json:"status" form:"status"`
		DueDate     time.Time `json:"due_date" form:"due_date"`
	}

	TaskResponse struct {
		ID          uuid.UUID `json:"id"`
		Title       string    `json:"title"`
		Description string    `json:"description"`
		Status      string    `json:"status"`
		DueDate     time.Time `json:"due_date"`
		UserID      string    `json:"user_id"`
		User        string    `json:"user,omitempty"`
		CreatedAt   time.Time `json:"created_at"`
		UpdatedAt   time.Time `json:"updated_at"`
	}

	TaskPaginationResponse struct {
		Data []TaskResponse `json:"data"`
		PaginationResponse
	}

	GetAllTasksRepositoryResponse struct {
		Tasks []entity.Task
		PaginationResponse
	}

	TaskUpdateRequest struct {
		Title       string    `json:"title" form:"title"`
		Description string    `json:"description" form:"description"`
		Status      string    `json:"status" form:"status"`
		DueDate     time.Time `json:"due_date" form:"due_date"`
	}

	TaskUpdateResponse struct {
		ID          uuid.UUID `json:"id"`
		Title       string    `json:"title"`
		Description string    `json:"description"`
		Status      string    `json:"status"`
		DueDate     time.Time `json:"due_date"`
		UserID      string    `json:"user_id"`
		CreatedAt   time.Time `json:"created_at"`
		UpdatedAt   time.Time `json:"updated_at"`
	}
)
