package controller

import (
	"net/http"

	"github.com/Revprm/go-fp-pbkk/dto"
	"github.com/Revprm/go-fp-pbkk/service"
	"github.com/Revprm/go-fp-pbkk/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type (
	TaskController interface {
		CreateTask(ctx *gin.Context)
		GetTaskByID(ctx *gin.Context)
		UpdateTask(ctx *gin.Context)
		DeleteTask(ctx *gin.Context)
		GetTasksWithPagination(ctx *gin.Context)
	}

	taskController struct {
		taskService service.TaskService
	}
)

func NewTaskController(taskService service.TaskService) TaskController {
	return &taskController{
		taskService: taskService,
	}
}

func (ctrl *taskController) CreateTask(ctx *gin.Context) {
	userId := ctx.MustGet("user_id").(string)

	var req dto.TaskCreateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	task, err := ctrl.taskService.CreateTask(ctx.Request.Context(), req, userId)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_REGISTER_USER, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_CREATE_TASK, task)
	ctx.JSON(http.StatusCreated, res)
}

func (ctrl *taskController) GetTaskByID(ctx *gin.Context) {
	taskID, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_TASK, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	task, err := ctrl.taskService.GetTaskByID(ctx.Request.Context(), taskID)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_TASK, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_GET_TASK, task)
	ctx.JSON(http.StatusOK, res)
}

func (ctrl *taskController) UpdateTask(ctx *gin.Context) {
	var req dto.TaskUpdateRequest
	taskID, err := uuid.Parse(ctx.Param("id"))
	if err := ctx.ShouldBind(&req); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	userID := ctx.MustGet("user_id").(string)
	result, err := ctrl.taskService.UpdateTask(ctx.Request.Context(), taskID, userID, req)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_UPDATE_TASK, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_UPDATE_TASK, result)
	ctx.JSON(http.StatusOK, res)
}

func (ctrl *taskController) DeleteTask(ctx *gin.Context) {
	taskID, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	userID := ctx.MustGet("user_id").(string)
	err = ctrl.taskService.DeleteTask(ctx.Request.Context(), taskID, userID)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_DELETE_TASK, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_DELETE_TASK, nil)
	ctx.JSON(http.StatusOK, res)
}

func (ctrl *taskController) GetTasksWithPagination(ctx *gin.Context) {
	userID := ctx.GetString("userID")
	var req dto.PaginationRequest
	if err := ctx.ShouldBind(&req); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	tasks, err := ctrl.taskService.GetTasksWithPagination(ctx.Request.Context(), userID, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := utils.Response{
		Status:  true,
		Message: dto.MESSAGE_SUCCESS_GET_LIST_USER,
		Data:    tasks.Data,
		Meta:    tasks.PaginationResponse,
	}

	ctx.JSON(http.StatusOK, response)
}
