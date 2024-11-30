package repository

import (
	"context"
	"math"

	"github.com/Revprm/go-fp-pbkk/dto"
	"github.com/Revprm/go-fp-pbkk/entity"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type (
	TaskRepository interface {
		CreateTask(ctx context.Context, tx *gorm.DB, task *entity.Task) (*entity.Task, error)
		GetTaskByID(ctx context.Context, tx *gorm.DB, taskID uuid.UUID) (*entity.Task, error)
		UpdateTask(ctx context.Context, tx *gorm.DB, taskID uuid.UUID, task *entity.Task) (*entity.Task, error)
		DeleteTask(ctx context.Context, tx *gorm.DB, taskID uuid.UUID) error
		GetTasksWithPagination(ctx context.Context, tx *gorm.DB, userID string, req dto.PaginationRequest) (*dto.GetAllTasksRepositoryResponse, error)
	}

	taskRepository struct {
		db *gorm.DB
	}
)

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &taskRepository{
		db: db,
	}
}

func (r *taskRepository) CreateTask(ctx context.Context, tx *gorm.DB, task *entity.Task) (*entity.Task, error) {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Create(task).Error; err != nil {
		return nil, dto.ErrCreateTask
	}

	return task, nil
}

func (r *taskRepository) GetTaskByID(ctx context.Context, tx *gorm.DB, taskID uuid.UUID) (*entity.Task, error) {
	if tx == nil {
		tx = r.db
	}

	var task entity.Task
	if err := tx.WithContext(ctx).Where("id = ?", taskID).Take(&task).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, dto.ErrTaskNotFound
		}
		return nil, dto.ErrGetTask
	}

	return &task, nil
}

func (r *taskRepository) UpdateTask(ctx context.Context, tx *gorm.DB, taskID uuid.UUID, task *entity.Task) (*entity.Task, error) {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Model(&entity.Task{}).Where("id = ?", taskID).Updates(task).Error; err != nil {
		return nil, dto.ErrUpdateTask
	}

	var updatedTask entity.Task
	if err := tx.WithContext(ctx).Where("id = ?", taskID).First(&updatedTask).Error; err != nil {
		return nil, dto.ErrTaskNotFound
	}

	return &updatedTask, nil
}

func (r *taskRepository) DeleteTask(ctx context.Context, tx *gorm.DB, taskID uuid.UUID) error {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Delete(&entity.Task{}, "id = ?", taskID).Error; err != nil {
		return dto.ErrDeleteTask
	}

	return nil
}

func (r *taskRepository) GetTasksWithPagination(ctx context.Context, tx *gorm.DB, userID string, req dto.PaginationRequest) (*dto.GetAllTasksRepositoryResponse, error) {
	if tx == nil {
		tx = r.db
	}

	var tasks []entity.Task
	var err error
	var count int64

	if req.PerPage == 0 {
		req.PerPage = 10
	}

	if req.Page == 0 {
		req.Page = 1
	}

	tx = tx.WithContext(ctx).Model(&entity.Task{})
	if userID != "" {
		tx = tx.Where("user_id = ?", userID)
	}

	if err := tx.Count(&count).Error; err != nil {
		return nil, dto.ErrGetTasks
	}

	if err := tx.Scopes(Paginate(req.Page, req.PerPage)).Find(&tasks).Error; err != nil {
		return nil, dto.ErrGetTasks
	}

	totalPage := int64(math.Ceil(float64(count) / float64(req.PerPage)))

	return &dto.GetAllTasksRepositoryResponse{
		Tasks: tasks,
		PaginationResponse: dto.PaginationResponse{
			Page:    req.Page,
			PerPage: req.PerPage,
			Count:   count,
			MaxPage: totalPage,
		},
	}, err
}
